# EVELib/SSO

EVE Online SSO Client functionality for Go.

**PLEASE NOTE**
1. Scopes passed are not checked to be real scopes at this time, instead leaving it up to the applications that utilize this package.
2. You don't need to space seperate scopes in `Client.Login` if you pass them as individual strings in the `scopes` variadic argument.
3. There may be some bugs, as I've not fully completed testing yet. Please keep this and mind, and consider submitting a new issue if you find a bug.
4. API is not yet guaranteed to be stable yet.
