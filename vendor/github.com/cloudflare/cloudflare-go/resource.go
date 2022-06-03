package cloudflare

// RouteLevel holds the "level" where the resource resides.
type RouteLevel string

const (
	AccountRouteLevel RouteLevel = "accounts"
	ZoneRouteLevel    RouteLevel = "zones"
	UserRouteLevel    RouteLevel = "user"
)

// ResourceContainer defines an API resource you wish to target. Should not be
// used directly, use `UserIdentifier`, `ZoneIdentifier` and `AccountIdentifier`
// instead.
type ResourceContainer struct {
	Level      RouteLevel
	Identifier string
}

// ResourceIdentifier returns a generic *ResourceContainer.
func ResourceIdentifier(id string) *ResourceContainer {
	return &ResourceContainer{
		Identifier: id,
	}
}

// UserIdentifier returns a user level *ResourceContainer.
func UserIdentifier(id string) *ResourceContainer {
	return &ResourceContainer{
		Level:      UserRouteLevel,
		Identifier: id,
	}
}

// ZoneIdentifier returns a zone level *ResourceContainer.
func ZoneIdentifier(id string) *ResourceContainer {
	return &ResourceContainer{
		Level:      ZoneRouteLevel,
		Identifier: id,
	}
}

// AccountIdentifier returns an account level *ResourceContainer.
func AccountIdentifier(id string) *ResourceContainer {
	return &ResourceContainer{
		Level:      AccountRouteLevel,
		Identifier: id,
	}
}
