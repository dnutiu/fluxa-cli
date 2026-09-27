package fluxa

import (
	"fmt"
	"github.com/dnutiu/fluxa-cli/internal/domain"
)

func collectionPath(entityID domain.ID, resource string) string {
	return fmt.Sprintf("/api/v1/entities/%d/%s", entityID, resource)
}

func itemPath(entityID domain.ID, resource string, recordID domain.ID) string {
	return fmt.Sprintf("%s/%d", collectionPath(entityID, resource), recordID)
}
