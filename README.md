# gorddo
Resilient Distributed Datasets. In Go.

## Usage

```go
import "github.com/taterbase/gorddo"

lines, err := gorrdo.TextFile("./file")
lines.Filter()
lines.Map()
lines.Collect()
```
