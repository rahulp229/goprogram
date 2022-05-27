package main

import (
	"fmt"
	"strings"
)

func main(){
	a := [5]int{1, 2, 3, 4, 5}
	s := a[:4] // output : {1,2,3}
	s[0] = 19 // output : {10,2,3}
	s = append(s,6) // output : {10,2,3,6}
	s = append(s,7) // output : {10,2,3,6,7}
	s = append(s,8)// output : {10,2,3,6,7,8}
	s[0] = 11 // output : {11,2,3,6,7,8}
	fmt.Println(a, s)// output :a={1, 2, 3, 4, 5} , s ={11,2,3,6,7,8}
	fmt.Println(cap(s), len(s)) //output : cap=nil, len=6

	sss := "https://report-check-objects-us-west-2-092721654985.s3.amazonaws.com/pdfs/23/3345322523.pdf?AWSAccessKeyId=ASIARLFVEFTEXE3BUOMY&Signature=MfsQdAkHdCAHCJYNVAkgxU0RkUQ%3D&x-amz-security-token=IQoJb3JpZ2luX2VjEBEaCXVzLXdlc3QtMiJHMEUCIQDKBghLJC3EAjkexbngp01%2BGKKypuhQZ5MdjWZeTXrvewIgU%2BxQUVC6nNzE496QKb1K2f%2BTxDUhJItQTNZmu6dtKCUqwwIIiv%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FARADGgwwOTI3MjE2NTQ5ODUiDDASSVadD9Dh9woooSqXAqHZTOhHddMT3P4c88lPwEwbgOau1p12xPDKrTpl0yLbK1xKElbaimQhL3mPFmc7khq0fN%2Bw8i8A4m7fnBb%2Bz9SjOAmtYGozhjpIMdoP8NnsN%2Fj5RG9HAhxRI94gVLmjftOs%2BA%2F5MX4qzwvGAMRsK3mZttqr%2F61S1q3fjrgyRWuzgNPG8j2zm1K16PY%2FGjauwd4b0Zf2qRIszXMLm%2Bea%2BHvMkh0yqBxyajtCZFBu%2Bgc9R4HFrBdio68UKQcL50UOaEirVnM27wjlSeUO7zJoqGeZfET4L2aZleuQWlVei1VGyhntD1Y4COiPs9zi5bic4AO4IDb4466kfKy8Cwlvq932Ud5vq7y%2BpNqib%2B3P1aA7vT00JejdhTDe5%2FqKBjqaAclQXOmyM%2FcoSQIEFTFRbUziwi0Lqry2PJYNO5HToXZf%2FKIRKfZhuQ5UnA%2B7ioViYpLGSnMr7%2Bm47dQcKXw1Njub0XxubqPs4dS9sZokmg5bmBbUVZfJ2YxlQvfp7jpbWiYxsqdhMiqytSBJiDnZuzV3hqRVWUlwlZfj33eTHFZ%2F9uS3tS79inSvxLEcpheTiB%2F4%2B1yg6MtvpC0%3D&Expires=1633600357"
	var sample = strings.Split(sss,"?")
	ss := sample[0]
	s1 := strings.Split(ss, ".")

	fmt.Println(">>>>",s1[len(s1)-1])
	fmt.Println(s)

}


