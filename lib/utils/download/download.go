// downlaoder package provides a single function Download
// which perform multiple downloads given a base url and
// multiple resources to download from that base url,
// downloads are done cuncurently, for each resource
// to download a new http.client is created as well as a
// new goroutune and returns channel of errors if any
// otherwise nil is returned in case of success (non error)

package download

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/sonic/lib/cmdui"
	"github.com/sonic/lib/errors"
)

var Silent bool = false

func Many(baseurl, saveto string, pkgs []string) []error {
	var (
		wg   errgroup.Group
		errs []error = []error{}
	)

	for _, pkgname := range pkgs {
		pkg := pkgname

		wg.Go(func() error {
			err := Single(baseurl, saveto, pkg)
			if err != nil {
				return errors.New(err.Error() + " : " + pkgname)
			}

			return nil
		})
	}

	if err := wg.Wait(); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func Single(baseurl, saveto, pkgname string) error {
	client := clientInit()

	resp, err := client.Get(baseurl + pkgname)
	if err != nil {
		return errors.Wrap(err)()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(resp.Status)
	}

	f, err := os.Create(path.Join(saveto, pkgname))
	if err != nil {
		return errors.Wrap(err)()
	}
	defer f.Close()

	return copy(resp.Body, f, resp.ContentLength, pkgname)
}

func copy(src io.Reader, dst io.Writer, srcsize int64, pkgname string) error {
	var (
		bufferSize int64 = 4096
		total      int64 = 0
	)

	buffer := make([]byte, bufferSize)
	body := io.LimitReader(src, srcsize)

	percent := srcsize / 100

	for {
		nreads, err := body.Read(buffer)
		if nreads > 0 {
			if err != nil && err != io.EOF {
				return errors.Wrap(err)()
			}

			total += int64(nreads)

			_, err = dst.Write(buffer[:nreads])
			if err != nil && err != io.EOF {
				return errors.Wrap(err)()
			}

			if !Silent {
				message := cmdui.NewMsg("Downloading " + pkgname)
				cmdui.ProgressPrinter(message, total, percent)
			}

			if total == srcsize {
				return nil
			}
		}

	}
}

func clientInit() http.Client {
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
	}
	return http.Client{Transport: transport}

}
