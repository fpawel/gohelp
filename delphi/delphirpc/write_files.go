package delphirpc

import (
	"github.com/fpawel/gohelp/must"
	"github.com/fpawel/gohelp/winapp"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

type SrcServices struct {
	Dir   string
	Types []reflect.Type
}
type SrcNotify struct {
	Dir         string
	Types       []NotifyServiceType
	PeerPackage string
}

func WriteSources(srv SrcServices, ntf SrcNotify) {
	if err := winapp.EnsuredDirectory(srv.Dir); err != nil {
		log.Fatal(err)
	}

	servicesSrc := NewServicesUnit(srv.Types)
	notifySvcSrc := NewNotifyServicesSrc(servicesSrc.TypesUnit, ntf.PeerPackage, ntf.Types)

	createFile := func(fileName string) *os.File {
		return must.Create(filepath.Join(srv.Dir, fileName))
	}

	file := createFile("services.pas")
	servicesSrc.WriteUnit(file)
	must.Close(file)

	file = createFile("server_data_types.pas")
	servicesSrc.TypesUnit.WriteUnit(file)
	must.Close(file)

	file = createFile("notify_services.pas")
	notifySvcSrc.WriteUnit(file)
	must.Close(file)

	file = must.Create(filepath.Join(ntf.Dir, "api_generated.go"))
	notifySvcSrc.WriteGoFile(file)
	must.Close(file)
}