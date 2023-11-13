GITHUB_URL=https://github.com/kubernetes/dashboard/releases
VERSION_KUBE_DASHBOARD=$(curl -I -L -s -S ${GITHUB_URL}/latest -o /dev/null | sed -e 's|.*/||')