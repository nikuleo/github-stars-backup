
## Github 收藏的仓库备份脚本
该脚本能够备份你 github 中 star 的仓库并以 YAML 文件的形式保存。
保存格式如下： 
```YAML
- repo name: DRQN_Stock_Trading
  author: conditionWang
  author's github: https://github.com/conditionWang
  url: https://github.com/conditionWang/DRQN_Stock_Trading
  description: ""
  language: Python
  tags: drqn, lstm, reinforcement-learning, stock-trading
- repo name: podcast
  author: bumingbaipod
  author's github: https://github.com/bumingbaipod
  url: https://github.com/bumingbaipod/podcast
  description: ""
  language: ""
  tags: ""
```

## 依赖
该脚本使用 `gopkg.in/yaml.v2` 包来序列化为 yaml 结构。

## 如何使用
### 命令参数：
```bash
-u [你的 github 用户名]
-p [你想要保存的文件路径]
-m [你收藏仓库的总数量]
```

### 使用案例：
```bash
go run . -u=nikusaikou -m=141
```