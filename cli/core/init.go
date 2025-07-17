package core

import (
	"fmt"
	"os"
	"path/filepath"
)
//folder
const TraceOpsDir = ".trace"
const ObjectFolder = "object"
//files
const ConfigFile = "config.json"
//files inside the object
// const VersionMeta = "version.meta.json"
// const VersionStruct = "files.json"
// const HashFile = "version.hash"

//optional
// const RefsFolder = "refs"
// const HeadFile = "HEAD"

func InitRepo() error {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("cannot get current directory: %w", err)
	}

	tracePath := filepath.Join(cwd, TraceOpsDir)

	// Check if the .trace directory already exists
	if _, err := os.Stat(tracePath); err == nil {
		return fmt.Errorf("traceops repo already initialized in %s", tracePath)
	}

	// Create .trace directory
	if err := os.MkdirAll(tracePath, 0755); err != nil {
		return fmt.Errorf("failed to create %s: %w", tracePath, err)
	}

	//Create object folder
	objectPath := filepath.Join(tracePath,ObjectFolder)
	if err := os.MkdirAll(objectPath,0755); err != nil{
		return fmt.Errorf("failed to create %s: %w", objectPath, err)
	}

	// Create config file
	configPath := filepath.Join(tracePath, ConfigFile)
	configJSON := `{
	  "user": {
	    "name": "",
	    "email": ""
	  },
	  "repos": []
	}`
	if err := os.WriteFile(configPath, []byte(configJSON), 0644); err != nil {
		return fmt.Errorf("failed to write config.json: %w", err)
	}

	fmt.Printf("Initialized empty TraceOps repo in %s\n", tracePath)
	return nil
}


// //create a file hash inside .trace/object
	// versionHashPath := filepath.Join(objectPath,HashFile )
	// if err := os.WriteFile(versionHashPath, []byte("ref: versions/v1.0\n"), 0644); err != nil {
	// 	return fmt.Errorf("failed to write HASH: %w", err)
	// }

	// versionHashPath := filepath.Join(objectPath,HashFile )
	// if err := os.WriteFile(versionHashPath, []byte("ref: versions/v1.0\n"), 0644); err != nil {
	// 	return fmt.Errorf("failed to write HASH: %w", err)
	// }



// //Create reference folder
	// refsPath := filepath.Join(tracePath,RefsFolder)
	// if err := os.MkdirAll(refsPath,0755); err != nil{
	// 	return fmt.Errorf("failed to create %s: %w", refsPath, err)
	// }

	// // Create HEAD file
	// headPath := filepath.Join(tracePath, HeadFile)
	// if err := os.WriteFile(headPath, []byte("ref: versions/v1.0\n"), 0644); err != nil {
	// 	return fmt.Errorf("failed to write HEAD: %w", err)
	// }