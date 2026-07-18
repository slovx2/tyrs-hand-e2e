---
name: issue-handler
description: Implement a GitHub issue end to end with tests and a pull request.
---

# Issue handler

1. 使用 GitHub 工具读取当前 Issue，确认验收条件。
2. 检查仓库现有代码与测试，只做 Issue 要求的最小改动。
3. 为新行为添加或更新自动化测试，并运行 `go test ./...`。
4. 使用本地 Git 工具检查状态并提交改动，提交信息使用英文 Conventional Commit。
5. 使用 `git.publish_branch` 发布当前托管分支。
6. 使用 GitHub 工具创建 Pull Request，标题清楚，正文包含 `Closes #<issue>` 和测试结果。
7. 在原 Issue 评论 PR 链接、实现摘要和测试结果。

禁止执行 merge、关闭 Issue/PR、强制推送或修改当前工作项以外的资源。
