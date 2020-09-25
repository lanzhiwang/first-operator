library "alauda-cicd"
def language = "golang"
AlaudaPipeline {
    config = [
        agent: 'golang-1.13',
        folder: '.',
        chart: [
            [
                chart: "tdsql",
                component: "first-operator",
                pipeline: "first-chart",
                branch:  "master",
            ],
        ],
        scm: [
            credentials: 'devops-alauda-gitlab'
        ],
        docker: [
            repository: "tdsql/imoocpod-operator",
            credentials: "tdsql-harbor-b",
            context: ".",
            dockerfile: "build/Dockerfile",
            enabled: true,
            armBuild: false,
        ],
        sonar: [
            binding: "sonarqube",
            enabled: false,
        ],
    ]
    env = [
        GO111MODULE: "on",
        GOPROXY: "https://athens.alauda.cn,https://goproxy.cn,direct",
        CGO_ENABLED: "0",
        GOOS: "linux",
    ]
    steps = [
        [
            name: "Test",
            container: "golang",
            commands: ["make test"]
        ],
        [
            name: "Build",
            container: "golang",
            commands: ["make build"]
        ],
    ]
}
