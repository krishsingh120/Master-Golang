package auth

import "fmt"

/*

- is function ka name abhi small letter se hai to ye sirf is auth package ke andar use kr sakte ho -> small mtlb private package.
- if agr is first letter capital krodo to ye khi bhi use ho sakta hai.
- koi bhi entity ya function ya variable bhar use krna hai us package ke to first letter capital.
- Third party packages in go:
    - pkg go site -> similar to npm
		- installation: go get <Package_URL>
		- file:
		   - go.mod => package.json
			 - go.sum => package-lock.json
			 - go mod tidy => npm install  -> fix everything upp to date.

*/

func LoginWithCredentials(username string, password string) {
	fmt.Println("User login with username and password:", username, password)
}
