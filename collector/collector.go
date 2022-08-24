
package collector

import (
	"github.com/erictsai1107/nutanix_exporter/nutanix"
	"strings"
)

var nutanixApi *nutanix.Nutanix

func normalizeFQN(fqn string) string {
	var _fqn string = fqn
	_fqn = strings.Replace(_fqn, ".", "_", -1)
	_fqn = strings.Replace(_fqn, "-", "_", -1)
	_fqn = strings.ToLower(_fqn)

	return _fqn
}
