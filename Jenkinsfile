library "alauda-cicd"
def language = "golang"
AlaudaPipeline{
    config = [
        agent: 'golang-1.12',
        folder: '.',
        chart: [
            chart: "test-chart",
            component: "test-chart",
        ],
        scm: [
            credentials: 'tdsql-zhihu'
        ],
        docker: [
            repository: "tdsql/imoocpod-operator",
            credentials: "tdsql-tdsql-harbor-b",
            context: ".",
            dockerfile: "build/Dockerfile",
        ],
        sonar: [
            binding: "tdsql-son"
        ],
    ]
    env = [
        GO111MODULE: "on",
        GOPROXY: "https://athens.alauda.cn,https://goproxy.cn,direct",
    ]
    steps = [
        [
            name: "Build",
            container: language,
            commands: [
                """make test
                make build""",
            ]
        ],
    ]
}
