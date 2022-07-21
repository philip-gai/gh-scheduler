# See here for image contents: https://github.com/github/debian-base
FROM ghcr.io/github/debian-base/debian-base:latest

# Install additional tools below

# Tools for the vscode-go extension
# https://github.com/golang/vscode-go/blob/HEAD/docs/tools.md
RUN go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest && \
    go install -v github.com/haya14busa/goplay/cmd/goplay@latest && \
    go install -v github.com/fatih/gomodifytags@latest && \
    go install -v github.com/josharian/impl@latest && \
    go install -v github.com/cweill/gotests/...@latest && \
    go install -v honnef.co/go/tools/cmd/staticcheck@latest
