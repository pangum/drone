package config

type Debug struct {
	// 应用名称
	Name string `default:"${NAME=${DRONE_STAGE_NAME}}" json:"name"`
	// 应用版本
	Version string `default:"${VERSION=${DRONE_TAG=${DRONE_COMMIT_BRANCH}}}" json:"version"`
	// 编译版本
	Build string `default:"${BUILD=${DRONE_BUILD_NUMBER}}" json:"build"`
	// 编译时间
	Complied string `default:"${TIMESTAMP=${DRONE_BUILD_STARTED}}" json:"complied"`
	// 分支版本
	Revision string `default:"${REVISION=${DRONE_COMMIT_SHA}}" json:"revision"`
	// 分支
	Branch string `default:"${BRANCH=${DRONE_COMMIT_BRANCH}}" json:"branch"`
}
