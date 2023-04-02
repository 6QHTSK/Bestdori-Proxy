# Bestdori-Proxy

提供一个BestdoriAPI的整合API，部分以反向代理的形式实现。

## API 文档

### 服务器 `server`

目前本API能够提供支持整合信息的服务器包括，各个API支持的服务器将会在后续给出

| `server`字段名 | 解释                     |
|-------------|------------------------|
| `bestdori`  | Bestdori自制谱面集          |
| `bandori`   | Bandori BanGDream官方谱面集 |
| `llsif`     | LoveLive SIF谱面集        |

### 谱面列表获取

#### ``GET /post/:server/list?offset=0&limit=3&username=merjoex``

- Server: `bestdori`,`bandori`

| 字段       | 解释     | 默认值  | 备注                |
|----------|--------|------|-------------------|
| offset   | 谱面列表偏移 | `0`  |                   |
| limit    | 单页谱面限制 | `20` | 取值范围 10 ~ 50      |
| username | 搜索用户名  | `""` | `bandori`服务器忽略此字段 |

#### 正常Response

```json
{
  "count": 30517,
  "list": [
    101987,
    101984,
    101982,
    101978,
    101977,
    101976,
    101975,
    101974,
    101970,
    101967
  ]
}
```

- count 总谱面数或`username`谱面数或官谱谱面数
- list 一个长度为`limit`的数组，其内容为包含的ID，当达到末尾时可能会少于limit

### 谱面信息获取

#### ``GET /post/:server/:postID/:method?diff=3``

- Server:`bestdori`,`bandori`
- 也可以通过 ``GET /post/:server/:postID?diff=3``以`full`的method获取谱面信息

| 字段     | 解释     | 默认值  | 备注                 |
|--------|--------|------|--------------------|
| postID | 谱面ID   | 必填   |                    |
| method | 结果返回方法 | `""` | 具体解释见下             |
| diff   | 谱面难度   | `3`  | `bestdori`服务器忽略此字段 |

- `method`字段

| `method`字段名 | 解释                      |
|-------------|-------------------------|
| `full`      | 返回Json为完整的谱面帖子信息        |
| `info`      | 在`full`方法基础上删去`chart`字段 |
| `chart`     | 返回Json仅包含`chart`字段      |

- `diff` 字段超出范围将强制修改为默认值

| server    | `diff`取值范围 | 默认值 |
|-----------|------------|-----|
| `bandori` | 0~4        | 3   |

#### 正常Response

```json
{
  "id": 128,
  "title": "六兆年と一夜物語",
  "artists": "Roselia",
  "username": "craftegg",
  "diff": 3,
  "rating": 29,
  "audioURL": "https://bestdori.com/assets/jp/sound/bgm128_rip/bgm128.mp3",
  "coverURL": "https://bestdori.com/assets/jp/musicjacket/musicjacket130_rip/assets-star-forassetbundle-startapp-musicjacket-musicjacket130-128_ichiyamonogatari-jacket.png",
  "time": 1528092000,
  "content": "六兆年と一夜物語",
  "chart": []
}
```

- chart字段略去，其为BestdoriV2格式
- Time 字段为发布时间，官谱为最早发布的服务器的时间，自制谱为自制谱面发布的时间
- 所有的字段，按照日服-国际服-台服-国服-韩服的顺序，选择最先一个不为null的展示

### 封面与音频的资源反代

#### 封面反代 ``GET /assets/:server/:postID/cover``

#### 音频反代 ``GET /assets/:server/:postID/audio``

- Server:`bestdori`,`bandori`,`llsif`

| 字段     | 解释   | 默认值 | 备注  |
|--------|------|-----|-----|
| postID | 谱面ID | 必填  |     |

### 错误Response

- 当发生错误时，Response的Status Code将不再是200，为错误码（400/404等）
- 返回Json含有err_code和err_msg，示例如下

```json
{
  "err_code": 211,
  "err_msg": "谱面未找到"
}
```

## 名词解释

- `post` 谱面帖子，一个谱面帖子包含info（信息部分）和chart（谱面部分）
- `postID` 谱面帖子ID，简称谱面ID
- `bestdori` Bestdori为自制谱面上传社区
- `bandori` BanGDream缩写，为官方谱面集
- `info` 帖子的各种信息，包括名称、艺术家等等
- `chart` Bestdori V2格式的谱面，官方谱面将被转化为自制谱面格式
- `diff` 难度，分为`easy`、`normal`、`hard`、`expert`、`special`五个难度
- `rating` 等级，目前等级跨度为5-35级
- `audio` 谱面音频
- `cover` 谱面封面