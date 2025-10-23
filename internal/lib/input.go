package lib

import (
	"bufio"
	"os"
	"strconv"
)

func Input(toType string)(any, error) {
	scanner := bufio.NewScanner(os.Stdin)
	switch toType {
		case "string":
		for scanner.Scan() {
			return scanner.Text(), nil
		}
		case "int":
		for scanner.Scan() {
			num,_ := strconv.Atoi(scanner.Text())
			return num, nil
		}
		case "bool":
		for scanner.Scan() {
			inp := scanner.Text()
			switch inp {
			case "y":
				return true, nil
			case "n":
				return false, nil
			default:
				res,_ := strconv.ParseBool(inp)
				return res, nil
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}