package relationships

import (
	uuid "github.com/gofrs/uuid"
	"gitlab.com/archstack/workspace-api/internal/platform/datastore"
	"gitlab.com/archstack/workspace-api/models"
)

// Relationships struct holds all the dependencies required for the relationships package. And exposes all services
// provided by this package as its methods
type Relationships struct {
	Repository RelationshipRepository
}

// NewService creates a new Relationships service
func NewService(datastore *datastore.Datastore) (*Relationships, error) {
	repo := RelationshipRepository{datastore}

	w := &Relationships{repo}

	return w, nil
}

// WorkspaceAndUser represents a many to many relationship between workspaces and users
type WorkspaceAndUser struct {
	tableName struct{} `pg:"workspaces_users"`

	WorkspaceID uuid.UUID `pg:",type:uuid,unique:idx_workspace_id_user_id"`
	UserID      uuid.UUID `pg:",type:uuid,unique:idx_workspace_id_user_id"`
}

// NewWorkspaceAndUser returns an instance of the WorkspaceAndUser many to many relationship
func NewWorkspaceAndUser(workspace *models.Workspace, user *models.User) WorkspaceAndUser {
	return WorkspaceAndUser{WorkspaceID: workspace.ID, UserID: user.ID}
}