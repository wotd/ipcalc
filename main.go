package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	ipAddressPtr := flag.String("a", "192.168.1.1", "IP Address")
	netMaskPtr := flag.String("n", "255.255.255.0", "Network Mask")

	flag.Parse()

	var ipIntArr, netIntArr, andIntRes, orIntRes []int

	ipIntArr = validateFlags(*ipAddressPtr)
	netIntArr = validateFlags(*netMaskPtr)

	for c := 0; c < 4; c++ {
		andIntRes = append(andIntRes, binaryAnd(ipIntArr[c], netIntArr[c]))
		orIntRes = append(orIntRes, binaryOR(ipIntArr[c], 255^netIntArr[c]))
	}
	hostMinArr := make([]int, len(andIntRes))
	copy(hostMinArr, andIntRes)
	hostMinArr[3] += 1

	hostMaxArr := make([]int, len(orIntRes))
	copy(hostMaxArr, orIntRes)
	hostMaxArr[3] -= 1

	fmt.Printf("Address:\t %v \t -> %v \n", *ipAddressPtr, strings.Join(binArr(ipIntArr), "."))
	fmt.Printf("Netmask:\t %v \t -> %v \n", *netMaskPtr, strings.Join(binArr(netIntArr), "."))
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Printf("Network:\t %v \t -> %v \n", strings.Join(intToString(andIntRes), "."), strings.Join(binArr(andIntRes), "."))
	fmt.Printf("Broadcast:\t %v \t -> %v \n", strings.Join(intToString(orIntRes), "."), strings.Join(binArr(orIntRes), "."))
	fmt.Printf("Host min:\t %v \t -> %v \n", strings.Join(intToString(hostMinArr), "."), strings.Join(binArr(hostMinArr), "."))
	fmt.Printf("Host max:\t %v \t -> %v \n", strings.Join(intToString(hostMaxArr), "."), strings.Join(binArr(hostMaxArr), "."))

}

func binaryAnd(a, b int) int {
	r := a & b
	return r
}

func binaryOR(a, b int) int {
	r := a | b
	return r
}

func binArr(arr []int) []string {
	var r []string

	for c := 0; c < 4; c++ {
		tmp := strconv.FormatInt(int64(arr[c]), 2)
		rBin := ""
		for i := 0; i < 8-len(tmp); i++ {
			rBin += "0"
		}
		rBin += tmp
		r = append(r, rBin)
	}

	return r
}

func intToString(a []int) []string {
	var tmp []string
	for _, f := range a {
		tmp = append(tmp, strconv.Itoa(f))
	}
	return tmp
}

func validateFlags(arg string) []int {

	tmp := strings.Split(arg, ".")
	var octet []int
	if len(tmp) != 4 {
		log.Fatal("Provide valid IP address / network mask")
	}

	for _, v := range tmp {
		o, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		octet = append(octet, o)
	}
	return octet
}
