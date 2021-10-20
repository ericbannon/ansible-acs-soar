package main

import (
        "log"
        "net/http"
        "time"
        "encoding/json"
        "io/ioutil"
        "github.com/gorilla/mux"
)

type SRAlert struct {
	Alert struct {
		ID     string `json:"id"`
		Policy struct {
			ID          string   `json:"id"`
			Name        string   `json:"name"`
			Description string   `json:"description"`
			Rationale   string   `json:"rationale"`
			Remediation string   `json:"remediation"`
			Disabled    bool     `json:"disabled"`
			Categories  []string `json:"categories"`
			Fields      struct {
				ImageName struct {
					Registry string `json:"registry"`
					Remote   string `json:"remote"`
					Tag      string `json:"tag"`
				} `json:"imageName"`
				ImageAgeDays string `json:"imageAgeDays"`
				LineRule     struct {
					Instruction string `json:"instruction"`
					Value       string `json:"value"`
				} `json:"lineRule"`
				Cvss struct {
					Op    string `json:"op"`
					Value int    `json:"value"`
				} `json:"cvss"`
				Cve       string `json:"cve"`
				Component struct {
					Name    string `json:"name"`
					Version string `json:"version"`
				} `json:"component"`
				ScanAgeDays  string `json:"scanAgeDays"`
				NoScanExists bool   `json:"noScanExists"`
				Env          struct {
					Key          string `json:"key"`
					Value        string `json:"value"`
					EnvVarSource string `json:"envVarSource"`
				} `json:"env"`
				Command      string `json:"command"`
				Args         string `json:"args"`
				Directory    string `json:"directory"`
				User         string `json:"user"`
				VolumePolicy struct {
					Name        string `json:"name"`
					Source      string `json:"source"`
					Destination string `json:"destination"`
					ReadOnly    bool   `json:"readOnly"`
					Type        string `json:"type"`
				} `json:"volumePolicy"`
				PortPolicy struct {
					Port     int    `json:"port"`
					Protocol string `json:"protocol"`
				} `json:"portPolicy"`
				RequiredLabel struct {
					Key          string `json:"key"`
					Value        string `json:"value"`
					EnvVarSource string `json:"envVarSource"`
				} `json:"requiredLabel"`
				RequiredAnnotation struct {
					Key          string `json:"key"`
					Value        string `json:"value"`
					EnvVarSource string `json:"envVarSource"`
				} `json:"requiredAnnotation"`
				DisallowedAnnotation struct {
					Key          string `json:"key"`
					Value        string `json:"value"`
					EnvVarSource string `json:"envVarSource"`
				} `json:"disallowedAnnotation"`
				Privileged              bool     `json:"privileged"`
				DropCapabilities        []string `json:"dropCapabilities"`
				AddCapabilities         []string `json:"addCapabilities"`
				ContainerResourcePolicy struct {
					CPUResourceRequest struct {
						Op    string `json:"op"`
						Value int    `json:"value"`
					} `json:"cpuResourceRequest"`
					CPUResourceLimit struct {
						Op    string `json:"op"`
						Value int    `json:"value"`
					} `json:"cpuResourceLimit"`
					MemoryResourceRequest struct {
						Op    string `json:"op"`
						Value int    `json:"value"`
					} `json:"memoryResourceRequest"`
					MemoryResourceLimit struct {
						Op    string `json:"op"`
						Value int    `json:"value"`
					} `json:"memoryResourceLimit"`
				} `json:"containerResourcePolicy"`
				ProcessPolicy struct {
					Name     string `json:"name"`
					Args     string `json:"args"`
					Ancestor string `json:"ancestor"`
					UID      string `json:"uid"`
				} `json:"processPolicy"`
				ReadOnlyRootFs     bool   `json:"readOnlyRootFs"`
				FixedBy            string `json:"fixedBy"`
				PortExposurePolicy struct {
					ExposureLevels []string `json:"exposureLevels"`
				} `json:"portExposurePolicy"`
				PermissionPolicy struct {
					PermissionLevel string `json:"permissionLevel"`
				} `json:"permissionPolicy"`
				HostMountPolicy struct {
					ReadOnly bool `json:"readOnly"`
				} `json:"hostMountPolicy"`
				WhitelistEnabled   bool `json:"whitelistEnabled"`
				RequiredImageLabel struct {
					Key          string `json:"key"`
					Value        string `json:"value"`
					EnvVarSource string `json:"envVarSource"`
				} `json:"requiredImageLabel"`
				DisallowedImageLabel struct {
					Key          string `json:"key"`
					Value        string `json:"value"`
					EnvVarSource string `json:"envVarSource"`
				} `json:"disallowedImageLabel"`
			} `json:"fields"`
			LifecycleStages []string `json:"lifecycleStages"`
			Whitelists      []struct {
				Name       string `json:"name"`
				Deployment struct {
					Name  string `json:"name"`
					Scope struct {
						Cluster   string `json:"cluster"`
						Namespace string `json:"namespace"`
						Label     struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"label"`
					} `json:"scope"`
				} `json:"deployment"`
				Image struct {
					Name string `json:"name"`
				} `json:"image"`
				Expiration time.Time `json:"expiration"`
			} `json:"whitelists"`
			Scope []struct {
				Cluster   string `json:"cluster"`
				Namespace string `json:"namespace"`
				Label     struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"label"`
			} `json:"scope"`
			Severity           string    `json:"severity"`
			EnforcementActions []string  `json:"enforcementActions"`
			Notifiers          []string  `json:"notifiers"`
			LastUpdated        time.Time `json:"lastUpdated"`
			SORTName           string    `json:"SORTName"`
			SORTLifecycleStage string    `json:"SORTLifecycleStage"`
			SORTEnforcement    bool      `json:"SORTEnforcement"`
			PolicyVersion      string    `json:"policyVersion"`
			PolicySections     []struct {
				SectionName  string `json:"sectionName"`
				PolicyGroups []struct {
					FieldName       string `json:"fieldName"`
					BooleanOperator string `json:"booleanOperator"`
					Negate          bool   `json:"negate"`
					Values          []struct {
						Value string `json:"value"`
					} `json:"values"`
				} `json:"policyGroups"`
			} `json:"policySections"`
		} `json:"policy"`
		LifecycleStage string `json:"lifecycleStage"`
		Deployment     struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Namespace   string `json:"namespace"`
			NamespaceID string `json:"namespaceId"`
			Labels      struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"labels"`
			ClusterID   string `json:"clusterId"`
			ClusterName string `json:"clusterName"`
			Containers  []struct {
				Image struct {
					ID   string `json:"id"`
					Name struct {
						Registry string `json:"registry"`
						Remote   string `json:"remote"`
						Tag      string `json:"tag"`
						FullName string `json:"fullName"`
					} `json:"name"`
					NotPullable bool `json:"notPullable"`
				} `json:"image"`
				Name string `json:"name"`
			} `json:"containers"`
			Annotations struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"annotations"`
			Inactive bool `json:"inactive"`
		} `json:"deployment"`
		Violations []struct {
			Message             string `json:"message"`
			DEPRECATEDProcesses []struct {
				ID            string `json:"id"`
				DeploymentID  string `json:"deploymentId"`
				ContainerName string `json:"containerName"`
				PodID         string `json:"podId"`
				PodUID        string `json:"podUid"`
				Signal        struct {
					ID           string    `json:"id"`
					ContainerID  string    `json:"containerId"`
					Time         time.Time `json:"time"`
					Name         string    `json:"name"`
					Args         string    `json:"args"`
					ExecFilePath string    `json:"execFilePath"`
					Pid          int       `json:"pid"`
					UID          int       `json:"uid"`
					Gid          int       `json:"gid"`
					Lineage      []string  `json:"lineage"`
					Scraped      bool      `json:"scraped"`
					LineageInfo  []struct {
						ParentUID          int    `json:"parentUid"`
						ParentExecFilePath string `json:"parentExecFilePath"`
					} `json:"lineageInfo"`
				} `json:"signal"`
				ClusterID          string    `json:"clusterId"`
				Namespace          string    `json:"namespace"`
				ContainerStartTime time.Time `json:"containerStartTime"`
			} `json:"DEPRECATEDProcesses"`
		} `json:"violations"`
		ProcessViolation struct {
			Message   string `json:"message"`
			Processes []struct {
				ID            string `json:"id"`
				DeploymentID  string `json:"deploymentId"`
				ContainerName string `json:"containerName"`
				PodID         string `json:"podId"`
				PodUID        string `json:"podUid"`
				Signal        struct {
					ID           string    `json:"id"`
					ContainerID  string    `json:"containerId"`
					Time         time.Time `json:"time"`
					Name         string    `json:"name"`
					Args         string    `json:"args"`
					ExecFilePath string    `json:"execFilePath"`
					Pid          int       `json:"pid"`
					UID          int       `json:"uid"`
					Gid          int       `json:"gid"`
					Lineage      []string  `json:"lineage"`
					Scraped      bool      `json:"scraped"`
					LineageInfo  []struct {
						ParentUID          int    `json:"parentUid"`
						ParentExecFilePath string `json:"parentExecFilePath"`
					} `json:"lineageInfo"`
				} `json:"signal"`
				ClusterID          string    `json:"clusterId"`
				Namespace          string    `json:"namespace"`
				ContainerStartTime time.Time `json:"containerStartTime"`
			} `json:"processes"`
		} `json:"processViolation"`
		Enforcement struct {
			Action  string `json:"action"`
			Message string `json:"message"`
		} `json:"enforcement"`
		Time          time.Time `json:"time"`
		FirstOccurred time.Time `json:"firstOccurred"`
		State         string    `json:"state"`
		SnoozeTill    time.Time `json:"snoozeTill"`
		Tags          []string  `json:"tags"`
	} `json:"alert"`
}

type ViolationResponse struct {
    ViolationID string
    Name string
}

type PlaybookMap []struct {
	Name     string `json:"Name"`
	Policy   string `json:"Policy"`
	Playbook string `json:"Playbook"`
}


// global tracker for last update initialized to zero... a long time ago
var lastUpdateTimestamp time.Time = time.Time{}
var playbooks PlaybookMap

func readPlaybookConfig () {
    // first, check if there's been any change
    if (time.Since(lastUpdateTimestamp).Seconds() < 60) {
        log.Println("Config Map change too recent for update")
        return
    }

    // if changed, read values from config
    log.Println ("opening config:")

    file, _ := ioutil.ReadFile("/var/run/playbook.json")

    err := json.Unmarshal([]byte(file), &playbooks)
    if (err != nil) {
        log.Println(err)
        return
    }

    log.Println("found %s entries", len(playbooks))

    // update last timestamp
    lastUpdateTimestamp = time.Now()
    return
}

func HomeHandler (w http.ResponseWriter, r *http.Request) {
    var violation SRAlert

    // update the playbook configuration
    readPlaybookConfig()

    if err := json.NewDecoder(r.Body).Decode(&violation); err != nil {
        log.Println(err)
        http.Error(w, "Error decoding alert from ACS", http.StatusBadRequest)
        return
    }
    // dispatch the playbook
    for _, playbook := range playbooks {
        if (playbook.Policy == violation.Alert.Policy.Name) {
            log.Println("Dispatching playbook %s", playbook.Policy)
        }
    }

    // response
    response := ViolationResponse{ViolationID:"1", Name:"placeholder"}
    respJSON, err := json.Marshal(response)
    if (err != nil) {
        log.Println(err)
        http.Error(w, "Error building response object", http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write(respJSON)
}


func main() {
    // get initial config from configMap
    readPlaybookConfig()

    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler).Methods("POST")
    http.Handle("/", r)
    http.ListenAndServe(":8080", r)
}
