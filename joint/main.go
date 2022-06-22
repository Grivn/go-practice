package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	sub1()

	sub2()
}

var tagValueVariable = regexp.MustCompile(`^tag_values+\(([a-zA-Z0-9-_.]+),([a-zA-Z0-9-_.]+)(,[a-zA-Z0-9-_.=$|,]+)?\)$`)

func sub1() {
	raw := "tag_values($v123,_p$system_3$system_4.system)"
	names := tagValueVariable.FindStringSubmatch(raw)
	fmt.Println(len(names))
	fmt.Println(names)

	fmt.Println(len(strings.Split(raw, ",")))
	fmt.Println(strings.Split(raw, ","))

	fmt.Println(len(strings.Split(raw, ".")))
	fmt.Println(strings.Split(raw, "."))

	fn := func(r rune) bool {
		return r == ',' || r == '.'
	}

	fmt.Println(len(strings.FieldsFunc(raw, fn)))
	fmt.Println(strings.FieldsFunc(raw, fn))
}

func formatRaw(raw string) string {
	raw = strings.Replace(raw, " ", "", -1)
	raw = strings.Replace(raw, "\n", "", -1)
	return raw
}

func sub2() {
	raw := "tag$part_1$v123,_u$system_3$system_4.system$part_2"
	raw = formatRaw(raw)

	temp := make(map[string]string)
	temp["part_1"] = "_values("
	temp["v123"] = "xxx.sdk.span.client.rate,host"
	temp["system_3"] = "rl=ha"
	temp["system_4"] = "ve.fun"
	temp["part_2"] = ")"

	fmt.Printf("format raw: %s\n", raw)

	//fn := func(r rune) bool {
	//	return r == ',' || r == '.'
	//}
	//
	//items := strings.FieldsFunc(raw, fn)

	split1 := strings.Split(raw, ",")

	fmt.Println(len(split1))
	fmt.Println(split1)

	res := ""
	for k, item := range split1 {
		split2 := strings.Split(item, ".")
		fmt.Println(len(split2))
		fmt.Println(split2)
		res1 := ""
		for i, item2 := range split2 {
			variables := strings.Split(item2, "$")
			fmt.Printf("len %v, %v\n", len(variables), variables)

			res2 := ""
			for index, variable := range variables {
				if index == 0 {
					res2 += variable
					continue
				}
				res2 += temp[variable]
			}
			res1 += res2
			if i < len(split2)-1 {
				res1 += "."
			}
		}
		fmt.Println(res1)
		res += res1
		if k < len(split1)-1 {
			res += ","
		}
	}
	fmt.Println(res)
}

type accessor struct {
	content map[string]string
}

// NOTE: we may need to find a better way to deal with it.
func (acc *accessor) joint(raw string) string {
	// split by "("
	leftParenJoint := ""
	leftParenSplit := strings.Split(raw, "(")
	for lpIndex, lpItem := range leftParenSplit {
		// split by ")"
		rightParenJoint := ""
		rightParenSplit := strings.Split(lpItem, ")")
		for rpIndex, rpItem := range rightParenSplit {
			// split by ","
			commaJoint := ""
			commaSplit := strings.Split(rpItem, ",")
			for commaIndex, commaItem := range commaSplit {
				// split by "."
				dotJoint := ""
				dotSplit := strings.Split(commaItem, ".")
				for dotIndex, dotItem := range dotSplit {
					// split by "$"
					variableJoint := ""
					variableSplit := strings.Split(dotItem, "$")
					for variableIndex, variableItem := range variableSplit {
						if variableIndex == 0 {
							// the first item cannot be a variable, joint directly.
							variableJoint += variableItem
							continue
						}
						value := acc.content[variableItem]
						if value == "" {
							variableJoint += fmt.Sprintf("$%s", variableItem)
						} else {
							variableJoint += value
						}
					}
					dotJoint += variableJoint
					if dotIndex < len(dotSplit)-1 {
						dotJoint += "."
					}
				}
				commaJoint += dotJoint
				if commaIndex < len(commaSplit)-1 {
					commaJoint += ","
				}
			}
			rightParenJoint += commaJoint
			if rpIndex < len(rightParenSplit)-1 {
				rightParenJoint += ")"
			}
		}
		leftParenJoint += rightParenJoint
		if lpIndex < len(leftParenSplit)-1 {
			leftParenJoint += "("
		}
	}
	return leftParenJoint
}

func (acc *accessor) splitAndJoint(raw string, sepList []string, index int) string {
	// the joint string to return.
	// sepList [",", ".", "(", ")"]
	joint := ""

	// split by the sep in sep_list
	if index < len(sepList) {
		sep := sepList[index]
		split := strings.Split(raw, sep)
		for i, item := range split {
			joint += acc.splitAndJoint(item, sepList, index+1)
			if i < len(split)-1 {
				joint += sep
			}
		}
		return joint
	}

	// the sep_list is empty,
	// then joint the variables according to "$" sep.
	split := strings.Split(raw, "$")
	for i, item := range split {
		if i == 0 {
			// the first item cannot be a variable, joint directly.
			joint += item
			continue
		}
		value := acc.content[item]
		if value == "" {
			variable := "$" + item
			joint += variable
		} else {
			joint += value
		}
	}
	return joint
}
