package cloudflare

<<<<<<< HEAD
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
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
import "fmt"

// RouteLevel holds the "level" where the resource resides. Commonly used in
// routing configurations or builders.
type RouteLevel string

// ResourceType holds the type of the resource. This is similar to `RouteLevel`
// however this is the singular version of `RouteLevel` and isn't suitable for
// use in routing.
type ResourceType string

const (
	user    = "user"
	zone    = "zone"
	account = "account"

	zones    = zone + "s"
	accounts = account + "s"

	AccountRouteLevel RouteLevel = accounts
	ZoneRouteLevel    RouteLevel = zones
	UserRouteLevel    RouteLevel = user

	AccountType ResourceType = account
	ZoneType    ResourceType = zone
	UserType    ResourceType = user
)

// ResourceContainer defines an API resource you wish to target. Should not be
// used directly, use `UserIdentifier`, `ZoneIdentifier` and `AccountIdentifier`
// instead.
type ResourceContainer struct {
	Level      RouteLevel
	Identifier string
	Type       ResourceType
}

func (r RouteLevel) String() string {
	switch r {
	case AccountRouteLevel:
		return accounts
	case ZoneRouteLevel:
		return zones
	case UserRouteLevel:
		return user
	default:
		return "unknown"
	}
}

func (r ResourceType) String() string {
	switch r {
	case AccountType:
		return account
	case ZoneType:
		return zone
	case UserType:
		return user
	default:
		return "unknown"
	}
}

// Returns a URL fragment of the endpoint scoped by the container.
//
// For example, a zone identifier would have a fragment like "zones/foobar" while
// an account identifier would generate "accounts/foobar".
func (rc *ResourceContainer) URLFragment() string {
	if rc.Level == "" {
		return rc.Identifier
	}

	if rc.Level == UserRouteLevel {
		return user
	}

	return fmt.Sprintf("%s/%s", rc.Level, rc.Identifier)
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
		Type:       UserType,
	}
}

// ZoneIdentifier returns a zone level *ResourceContainer.
func ZoneIdentifier(id string) *ResourceContainer {
	return &ResourceContainer{
		Level:      ZoneRouteLevel,
		Identifier: id,
		Type:       ZoneType,
	}
}

// AccountIdentifier returns an account level *ResourceContainer.
func AccountIdentifier(id string) *ResourceContainer {
	return &ResourceContainer{
		Level:      AccountRouteLevel,
		Identifier: id,
		Type:       AccountType,
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}
}
