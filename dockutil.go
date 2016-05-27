package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	build := flag.Bool("build", false, "build a docker image, extract docker.name and docker.tag LABEL from Dockerfile")
	buildNoPull := flag.Bool("build-no-pull", false, "sme as build, but does not force pull from repo")
	run := flag.Bool("run", false, "run [docker args] a docker image, bind port from EXPOSED tag in Dockerfile, pass [args] to docker image (you can overrride CMD in Dockerfile)")
	exec := flag.Bool("exec", false, "exec a command in a running container (default to /bin/bash)")
	push := flag.Bool("push", false, "upload image to registry (-r is required)")
	rsc := flag.Bool("rsc", false, "remove stopped containers form docker engine ")
	rui := flag.Bool("rui", false, "remove untagged images")
	layer := flag.Bool("layer", false, "display docker layer sizes")

	flag.Parse()

	if len(os.Args) == 1 {
		printUsage()
	} else {
		fmt.Printf("the value for build is %t.\n", *build)
	}

}

func printUsage() {
	fmt.Println(`
Usage: $0 [-f docker_file] [-r docker_registry] [-t tag] [-n name] [ -o options] command     
command:    
	build = build a docker image, extract docker.name and docker.tag LABEL from Dockerfile   
	build-no-pull = sme as build, but does not force pull from repo   
	run [args]  = run a docker image, bind port from EXPOSED tag in Dockerfile, pass [args] to docker image (you can overrride CMD in Dockerfile)   
	exec [args] = exec a command in a running container (default to /bin/bash)   
	push = upload image to registry (-r is required)   
	rsc = remove stopped containers form docker engine   
	rui = remove untagged images   
	layer = display docker layer sizes 
 
options:    
	-f docker file path and name (default to a Dockerfile in current folder)   
	-t docker tag (default to LABEL docker.tag in Dockerfile)   
	-n docker name (default to LABEL docker.name in Dockerfile)   
	-r docker registry is required for push command   
	-o docker command options (rg. --privileged=true)   
    `)
}
