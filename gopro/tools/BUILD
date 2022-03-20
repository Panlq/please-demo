github_repo(
    name = "pleasings2",
    repo = "sagikazarmark/mypleasings",
    revision = f"{CONFIG.PLEASINGS2_VERSION}",
)

sh_cmd(
    name = "plz-tidy",
    cmd = [
        "$(out_exe :godeps) -dir third_party/go -clean -builtin -wollemi",
        "$(out_exe :wollemi) --log warn gofmt ...",
    ],
    deps = [
        ":godeps",
        ":wollemi",
    ],
)

HOME = "Users/jonpan"
remote_file(
    name = "godeps",
    # hashes = [
    #     "21dfcb8cacab410c351e3e57f3dbb79d77621da4bde42fad180f9d6ad82123c3",
    #     "16acdca286a85e73f2643ec2726f7fa749f0c657b92b4e67eeec43c275854b63",
    # ],
    # url = f"https://github.com/sagikazarmark/please-go-modules/releases/latest/download/godeps_{CONFIG.HOSTOS}_{CONFIG.HOSTARCH}.tar.gz",
    url = f"file:///{HOME}/go/bin/godeps",
    extract = False,
    exported_files = ["godeps"],
    binary = True,
)

remote_file(
    name = "wollemi",
    # url = f"https://github.com/tcncloud/wollemi/releases/download/{CONFIG.WOLLEMI_VERSION}/wollemi-{CONFIG.WOLLEMI_VERSION}-{CONFIG.HOSTOS}-{CONFIG.HOSTARCH}.tar.gz",
    url = f"file:///{HOME}/go/bin/wollemi",
    extract = False,
    exported_files = ["wollemi"],
    binary = True,
)