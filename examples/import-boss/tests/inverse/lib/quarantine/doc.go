// quarantine is inside the library, but should not import the private package. But it does!
package quarantine

import (
	_ "github.com/mochen302/gengox/examples/import-boss/tests/inverse/lib/private"
)
