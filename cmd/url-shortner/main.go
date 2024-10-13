package url_shortner

import (
	"RestApi/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoadConfig(cfg)

	fmt.Println(cfg)

}
