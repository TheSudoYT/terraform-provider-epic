package epic

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func transformSelection(d *schema.ResourceData, selected string) (string, error) {
	// Apply lower and upper case modifications.
	lower, lowerOk := d.Get("lower").(bool)
	upper, upperOk := d.Get("upper").(bool)
	if !lowerOk || !upperOk {
		return "", fmt.Errorf("expected boolean values for lower and upper")
	}

	if lower && !upper {
		selected = strings.ToLower(selected)
	} else if !lower && upper {
		selected = strings.ToUpper(selected)
	} else if lower && upper {
		// Do nothing, retain original case.
	} else {
		// Default behavior if both are false: convert to lower case.
		selected = strings.ToLower(selected)
	}
	return selected, nil
}
