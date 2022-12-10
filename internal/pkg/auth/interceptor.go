package auth

type Interceptor struct {
	accessibleRoles map[string][]string
}

func NewAuthInterceptor(accessibleRoles map[string][]string) Interceptor {
	return Interceptor{
		accessibleRoles: accessibleRoles,
	}
}
