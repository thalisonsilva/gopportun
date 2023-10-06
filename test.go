package gopportun_test

import (
	"testing"

	"github.com/thalisonsilva/gopportun.git/router"
)

func TestMain(m *testing.M) {
	router.Inicializa()
	m.Run()
}
