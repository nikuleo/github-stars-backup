
## Github 收藏的仓库备份脚本
该脚本能够备份你 github 中 star 的仓库并以 YAML 文件的形式保存。
保存格式如下： 
```YAML
- repo name: DRQN_Stock_Trading
  url: https://github.com/conditionWang/DRQN_Stock_Trading
  description: 'This is the code implementation of the paper "Financial Trading as
    a Game: A Deep Reinforcement Learning Approach".'
  author: conditionWang
  author's github: https://github.com/conditionWang
  language: Python
  tags: drqn, lstm, reinforcement-learning, stock-trading
- repo name: MES40
  url: https://github.com/KH40-khoast40/MES40
  description: MMD Extended Shader (MES40)
  author: KH40-khoast40
  author's github: https://github.com/KH40-khoast40
  language: HLSL
  tags: diva, fx, hlsl, materials, mikumikudance, mikumikueffect, mmd, mme, project,
    shader, shadow
```

## 依赖
该脚本使用 `gopkg.in/yaml.v2` 包来序列化为 yaml 结构。

## 如何使用
### 命令参数：
```bash
-u [你的 github 用户名]
-p [你想要保存的文件路径(可选项)]
-m [你收藏仓库的总数量]
-g [是否使用 goroutine 并行爬取(可选项，默认值为 true. 如果为 true 返回每一页之间结果无序)]
```

### 使用案例：
```bash
go run . -u=nikusaikou -p=stars -m=140 -g=false
```