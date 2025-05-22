
This is my work from [Roadmap.sh](https://roadmap.sh) on [second task](https://roadmap.sh/projects/github-user-activity)  

# Technology
I'm just using Go as programming language.  

# How to run
Run this code using command `go run . github-activity <username> <optional_github_event>`  

# Architecture
I just separate the codes by it's meaning like `types.go` for the structs declaration,  and `main.go` for logic and main executioner of this programs

# Example
cmd: `go run . github-activity Akihira77`
output:
```
2025/05/22 14:27:37 INFO Results Total=6 Type=ALL Data="[{Size:3 Type:CreateEvent Items:[{ID:49691105704 Type:CreateEvent Public:true Payload:{Ref: RefType:repository MasterBranch:main Description: PusherType:user} Repo:{ID:983181550 Name:roadmap-sh-Backend/github-user-activity URL:https://api.github.com/repos/roadmap-sh-Backend/github-user-activity} Actor:{ID:83736692 Login:Akihira77 GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/83736692? URL:https://api.github.com/users/Akihira77} Org:{ID:211434693 Login:roadmap-sh-Backend GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/211434693? URL:https://api.github.com/orgs/roadmap-sh-Backend} CreatedAt:2025-05-14 02:22:08 +0000 UTC} {ID:49643051216 Type:CreateEvent Public:true Payload:{Ref:master RefType:branch MasterBranch:main Description: PusherType:user} Repo:{ID:982508870 Name:roadmap-sh-Backend/task-tracker-cli URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli} Actor:{ID:83736692 Login:Akihira77 GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/83736692? URL:https://api.github.com/users/Akihira77} Org:{ID:211434693 Login:roadmap-sh-Backend GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/211434693? URL:https://api.github.com/orgs/roadmap-sh-Backend} CreatedAt:2025-05-13 02:04:10 +0000 UTC} {ID:49642996635 Type:CreateEvent Public:true Payload:{Ref: RefType:repository MasterBranch:main Description: PusherType:user} Repo:{ID:982508870 Name:roadmap-sh-Backend/task-tracker-cli URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli} Actor:{ID:83736692 Login:Akihira77 GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/83736692? URL:https://api.github.com/users/Akihira77} Org:{ID:211434693 Login:roadmap-sh-Backend GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/211434693? URL:https://api.github.com/orgs/roadmap-sh-Backend} CreatedAt:2025-05-13 02:01:28 +0000 UTC}]} {Size:3 Type:PushEvent Items:[{ID:49690963589 Type:PushEvent Public:true Payload:{PushID:24281818321 Size:1 DistinctSize:1 Ref:refs/heads/master Head:66ee8d138e1bd4cd372895ef0cfca589bb26b24a Before:b214e62ad5fe6034f0f9f5614f9fef7bd0a74325 Commits:[{Sha:66ee8d138e1bd4cd372895ef0cfca589bb26b24a Author:{Email:tik.dikawahyu123@gmail.com Name:Akihira77} Message:add README Distinct:true URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli/commits/66ee8d138e1bd4cd372895ef0cfca589bb26b24a}]} Repo:{ID:982508870 Name:roadmap-sh-Backend/task-tracker-cli URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli} Actor:{ID:83736692 Login:Akihira77 GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/83736692? URL:https://api.github.com/users/Akihira77} Org:{ID:211434693 Login:roadmap-sh-Backend GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/211434693? URL:https://api.github.com/orgs/roadmap-sh-Backend} CreatedAt:2025-05-14 02:14:39 +0000 UTC} {ID:49690780508 Type:PushEvent Public:true Payload:{PushID:24281725953 Size:1 DistinctSize:1 Ref:refs/heads/master Head:b214e62ad5fe6034f0f9f5614f9fef7bd0a74325 Before:4a6e00833bcbc993344189fb5864b3e8494df0e3 Commits:[{Sha:b214e62ad5fe6034f0f9f5614f9fef7bd0a74325 Author:{Email:tik.dikawahyu123@gmail.com Name:Akihira77} Message:feat: refactore code Distinct:true URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli/commits/b214e62ad5fe6034f0f9f5614f9fef7bd0a74325}]} Repo:{ID:982508870 Name:roadmap-sh-Backend/task-tracker-cli URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli} Actor:{ID:83736692 Login:Akihira77 GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/83736692? URL:https://api.github.com/users/Akihira77} Org:{ID:211434693 Login:roadmap-sh-Backend GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/211434693? URL:https://api.github.com/orgs/roadmap-sh-Backend} CreatedAt:2025-05-14 02:05:01 +0000 UTC} {ID:49645238825 Type:PushEvent Public:true Payload:{PushID:24259529995 Size:1 DistinctSize:1 Ref:refs/heads/master Head:4a6e00833bcbc993344189fb5864b3e8494df0e3 Before:fd14a59a8e559da1dc8575e865fd66855df9d205 Commits:[{Sha:4a6e00833bcbc993344189fb5864b3e8494df0e3 Author:{Email:tik.dikawahyu123@gmail.com Name:Akihira77} Message:feat: develop list, add, update, delete feature, and setup makefile Distinct:true URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli/commits/4a6e00833bcbc993344189fb5864b3e8494df0e3}]} Repo:{ID:982508870 Name:roadmap-sh-Backend/task-tracker-cli URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli} Actor:{ID:83736692 Login:Akihira77 GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/83736692? URL:https://api.github.com/users/Akihira77} Org:{ID:211434693 Login:roadmap-sh-Backend GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/211434693? URL:https://api.github.com/orgs/roadmap-sh-Backend} CreatedAt:2025-05-13 03:59:05 +0000 UTC}]}]
```
or using GPT for formatting the output becomes:
```
{
  "Total": 6,
  "Type": "ALL",
  "Data": [
    {
      "Size": 3,
      "Type": "CreateEvent",
      "Items": [
        {
          "ID": 49691105704,
          "Type": "CreateEvent",
          "Public": true,
          "Payload": {
            "Ref": "",
            "RefType": "repository",
            "MasterBranch": "main",
            "Description": "",
            "PusherType": "user"
          },
          "Repo": {
            "ID": 983181550,
            "Name": "roadmap-sh-Backend/github-user-activity",
            "URL": "https://api.github.com/repos/roadmap-sh-Backend/github-user-activity"
          },
          "Actor": {
            "ID": 83736692,
            "Login": "Akihira77",
            "GravatarID": "",
            "AvatarURL": "https://avatars.githubusercontent.com/u/83736692?",
            "URL": "https://api.github.com/users/Akihira77"
          },
          "Org": {
            "ID": 211434693,
            "Login": "roadmap-sh-Backend",
            "GravatarID": "",
            "AvatarURL": "https://avatars.githubusercontent.com/u/211434693?",
            "URL": "https://api.github.com/orgs/roadmap-sh-Backend"
          },
          "CreatedAt": "2025-05-14T02:22:08Z"
        },
        {
          "ID": 49643051216,
          "Type": "CreateEvent",
          "Public": true,
          "Payload": {
            "Ref": "master",
            "RefType": "branch",
            "MasterBranch": "main",
            "Description": "",
            "PusherType": "user"
          },
          "Repo": {
            "ID": 982508870,
            "Name": "roadmap-sh-Backend/task-tracker-cli",
            "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli"
          },
          "Actor": {
            "ID": 83736692,
            "Login": "Akihira77",
            "GravatarID": "",
            "AvatarURL": "https://avatars.githubusercontent.com/u/83736692?",
            "URL": "https://api.github.com/users/Akihira77"
          },
          "Org": {
            "ID": 211434693,
            "Login": "roadmap-sh-Backend",
            "GravatarID": "",
            "AvatarURL": "https://avatars.githubusercontent.com/u/211434693?",
            "URL": "https://api.github.com/orgs/roadmap-sh-Backend"
          },
          "CreatedAt": "2025-05-13T02:04:10Z"
        },
        {
          "ID": 49642996635,
          "Type": "CreateEvent",
          "Public": true,
          "Payload": {
            "Ref": "",
            "RefType": "repository",
            "MasterBranch": "main",
            "Description": "",
            "PusherType": "user"
          },
          "Repo": {
            "ID": 982508870,
            "Name": "roadmap-sh-Backend/task-tracker-cli",
            "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli"
          },
          "Actor": {
            "ID": 83736692,
            "Login": "Akihira77",
            "GravatarID": "",
            "AvatarURL": "https://avatars.githubusercontent.com/u/83736692?",
            "URL": "https://api.github.com/users/Akihira77"
          },
          "Org": {
            "ID": 211434693,
            "Login": "roadmap-sh-Backend",
            "GravatarID": "",
            "AvatarURL": "https://avatars.githubusercontent.com/u/211434693?",
            "URL": "https://api.github.com/orgs/roadmap-sh-Backend"
          },
          "CreatedAt": "2025-05-13T02:01:28Z"
        }
      ]
    },
    {
      "Size": 3,
      "Type": "PushEvent",
      "Items": [
        {
          "ID": 49690963589,
          "Type": "PushEvent",
          "Public": true,
          "Payload": {
            "PushID": 24281818321,
            "Size": 1,
            "DistinctSize": 1,
            "Ref": "refs/heads/master",
            "Head": "66ee8d138e1bd4cd372895ef0cfca589bb26b24a",
            "Before": "b214e62ad5fe6034f0f9f5614f9fef7bd0a74325",
            "Commits": [
              {
                "Sha": "66ee8d138e1bd4cd372895ef0cfca589bb26b24a",
                "Author": {
                  "Email": "tik.dikawahyu123@gmail.com",
                  "Name": "Akihira77"
                },
                "Message": "add README",
                "Distinct": true,
                "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli/commits/66ee8d138e1bd4cd372895ef0cfca589bb26b24a"
              }
            ]
          },
          "Repo": {
            "ID": 982508870,
            "Name": "roadmap-sh-Backend/task-tracker-cli",
            "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli"
          },
          "Actor": {
            "ID": 83736692,
            "Login": "Akihira77",
            "GravatarID": "",
            "AvatarURL": "https://avatars.githubusercontent.com/u/83736692?",
            "URL": "https://api.github.com/users/Akihira77"
          },
          "Org": {
            "ID": 211434693,
            "Login": "roadmap-sh-Backend",
            "GravatarID": "",
            "AvatarURL": "https://avatars.githubusercontent.com/u/211434693?",
            "URL": "https://api.github.com/orgs/roadmap-sh-Backend"
          },
          "CreatedAt": "2025-05-14T02:14:39Z"
        },
        {
          "ID": 49690780508,
          "Type": "PushEvent",
          "Public": true,
          "Payload": {
            "PushID": 24281725953,
            "Size": 1,
            "DistinctSize": 1,
            "Ref": "refs/heads/master",
            "Head": "b214e62ad5fe6034f0f9f5614f9fef7bd0a74325",
            "Before": "4a6e00833bcbc993344189fb5864b3e8494df0e3",
            "Commits": [
              {
                "Sha": "b214e62ad5fe6034f0f9f5614f9fef7bd0a74325",
                "Author": {
                  "Email": "tik.dikawahyu123@gmail.com",
                  "Name": "Akihira77"
                },
                "Message": "feat: refactore code",
                "Distinct": true,
                "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli/commits/b214e62ad5fe6034f0f9f5614f9fef7bd0a74325"
              }
            ]
          },
          "Repo": {
            "ID": 982508870,
            "Name": "roadmap-sh-Backend/task-tracker-cli",
            "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli"
          },
          "Actor": {
            "ID": 83736692,
            "Login": "Akihira77",
            "GravatarID": "",
            "AvatarURL": "https://avatars.githubusercontent.com/u/83736692?",
            "URL": "https://api.github.com/users/Akihira77"
          },
          "Org": {
            "ID": 211434693,
            "Login": "roadmap-sh-Backend",
            "GravatarID": "",
            "AvatarURL": "https://avatars.githubusercontent.com/u/211434693?",
            "URL": "https://api.github.com/orgs/roadmap-sh-Backend"
          },
          "CreatedAt": "2025-05-14T02:05:01Z"
        },
        {
          "ID": 49645238825,
          "Type": "PushEvent",
          "Public": true,
          "Payload": {
            "PushID": 24259529995,
            "Size": 1,
            "DistinctSize": 1,
            "Ref": "refs/heads/master",
            "Head": "4a6e00833bcbc993344189fb5864b3e8494df0e3",
            "Before": "fd14a59a8e559da1dc8575e865fd66855df9d205",
            "Commits": [
              {
                "Sha": "4a6e00833bcbc993344189fb5864b3e8494df0e3",
                "Author": {
                  "Email": "tik.dikawahyu123@gmail.com",
                  "Name": "Akihira77"
                },
                "Message": "feat: develop list, add, update, delete feature, and setup makefile",
                "Distinct": true,
                "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli/commits/4a6e00833bcbc993344189fb5864b3e8494df0e3"
              }
            ]
          },
          "Repo": {
            "ID": 982508870,
            "Name": "roadmap-sh-Backend/task-tracker-cli",
            "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli"
          },
          "Actor": {
            "ID": 83736692,
            "Login": "Akihira77",
            "GravatarID": "",
            "AvatarURL": "https://avatars.githubusercontent.com/u/83736692?",
            "URL": "https://api.github.com/users/Akihira77"
          },
          "Org": {
            "ID": 211434693,
            "Login": "roadmap-sh-Backend",
            "GravatarID": "",
            "AvatarURL": "https://avatars.githubusercontent.com/u/211434693?",
            "URL": "https://api.github.com/orgs/roadmap-sh-Backend"
          },
          "CreatedAt": "2025-05-13T03:59:05Z"
        }
      ]
    }
  ]
}
```
cmd: `go run . github-activity Akihira77 PushEvent`  
output: 
```
2025/05/22 14:20:57 Github Event PushEvent Type
2025/05/22 14:20:57 INFO Results Total=3 Type=PushEvent Data="[{Size:3 Type:PushEvent Items:[{ID:49690963589 Type:PushEvent Public:true Payload:{PushID:24281818321 Size:1 DistinctSize:1 Ref:refs/heads/master Head:66ee8d138e1bd4cd372895ef0cfca589bb26b24a Before:b214e62ad5fe6034f0f9f5614f9fef7bd0a74325 Commits:[{Sha:66ee8d138e1bd4cd372895ef0cfca589bb26b24a Author:{Email:tik.dikawahyu123@gmail.com Name:Akihira77} Message:add README Distinct:true URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli/commits/66ee8d138e1bd4cd372895ef0cfca589bb26b24a}]} Repo:{ID:982508870 Name:roadmap-sh-Backend/task-tracker-cli URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli} Actor:{ID:83736692 Login:Akihira77 GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/83736692? URL:https://api.github.com/users/Akihira77} Org:{ID:211434693 Login:roadmap-sh-Backend GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/211434693? URL:https://api.github.com/orgs/roadmap-sh-Backend} CreatedAt:2025-05-14 02:14:39 +0000 UTC} {ID:49690780508 Type:PushEvent Public:true Payload:{PushID:24281725953 Size:1 DistinctSize:1 Ref:refs/heads/master Head:b214e62ad5fe6034f0f9f5614f9fef7bd0a74325 Before:4a6e00833bcbc993344189fb5864b3e8494df0e3 Commits:[{Sha:b214e62ad5fe6034f0f9f5614f9fef7bd0a74325 Author:{Email:tik.dikawahyu123@gmail.com Name:Akihira77} Message:feat: refactore code Distinct:true URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli/commits/b214e62ad5fe6034f0f9f5614f9fef7bd0a74325}]} Repo:{ID:982508870 Name:roadmap-sh-Backend/task-tracker-cli URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli} Actor:{ID:83736692 Login:Akihira77 GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/83736692? URL:https://api.github.com/users/Akihira77} Org:{ID:211434693 Login:roadmap-sh-Backend GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/211434693? URL:https://api.github.com/orgs/roadmap-sh-Backend} CreatedAt:2025-05-14 02:05:01 +0000 UTC} {ID:49645238825 Type:PushEvent Public:true Payload:{PushID:24259529995 Size:1 DistinctSize:1 Ref:refs/heads/master Head:4a6e00833bcbc993344189fb5864b3e8494df0e3 Before:fd14a59a8e559da1dc8575e865fd66855df9d205 Commits:[{Sha:4a6e00833bcbc993344189fb5864b3e8494df0e3 Author:{Email:tik.dikawahyu123@gmail.com Name:Akihira77} Message:feat: develop list, add, update, delete feature, and setup makefile Distinct:true URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli/commits/4a6e00833bcbc993344189fb5864b3e8494df0e3}]} Repo:{ID:982508870 Name:roadmap-sh-Backend/task-tracker-cli URL:https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli} Actor:{ID:83736692 Login:Akihira77 GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/83736692? URL:https://api.github.com/users/Akihira77} Org:{ID:211434693 Login:roadmap-sh-Backend GravatarID: AvatarURL:https://avatars.githubusercontent.com/u/211434693? URL:https://api.github.com/orgs/roadmap-sh-Backend} CreatedAt:2025-05-13 03:59:05 +0000 UTC}]}]"
```
or using GPT for formatting the output becomes:
```
{
  "Total": 3,
  "Type": "PushEvent",
  "Data": [
    {
      "ID": 49690963589,
      "Type": "PushEvent",
      "Public": true,
      "Payload": {
        "PushID": 24281818321,
        "Size": 1,
        "DistinctSize": 1,
        "Ref": "refs/heads/master",
        "Head": "66ee8d138e1bd4cd372895ef0cfca589bb26b24a",
        "Before": "b214e62ad5fe6034f0f9f5614f9fef7bd0a74325",
        "Commits": [
          {
            "Sha": "66ee8d138e1bd4cd372895ef0cfca589bb26b24a",
            "Author": {
              "Email": "tik.dikawahyu123@gmail.com",
              "Name": "Akihira77"
            },
            "Message": "add README",
            "Distinct": true,
            "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli/commits/66ee8d138e1bd4cd372895ef0cfca589bb26b24a"
          }
        ]
      },
      "Repo": {
        "ID": 982508870,
        "Name": "roadmap-sh-Backend/task-tracker-cli",
        "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli"
      },
      "Actor": {
        "ID": 83736692,
        "Login": "Akihira77",
        "AvatarURL": "https://avatars.githubusercontent.com/u/83736692?",
        "URL": "https://api.github.com/users/Akihira77"
      },
      "Org": {
        "ID": 211434693,
        "Login": "roadmap-sh-Backend",
        "AvatarURL": "https://avatars.githubusercontent.com/u/211434693?",
        "URL": "https://api.github.com/orgs/roadmap-sh-Backend"
      },
      "CreatedAt": "2025-05-14T02:14:39Z"
    },
    {
      "ID": 49690780508,
      "Type": "PushEvent",
      "Public": true,
      "Payload": {
        "PushID": 24281725953,
        "Size": 1,
        "DistinctSize": 1,
        "Ref": "refs/heads/master",
        "Head": "b214e62ad5fe6034f0f9f5614f9fef7bd0a74325",
        "Before": "4a6e00833bcbc993344189fb5864b3e8494df0e3",
        "Commits": [
          {
            "Sha": "b214e62ad5fe6034f0f9f5614f9fef7bd0a74325",
            "Author": {
              "Email": "tik.dikawahyu123@gmail.com",
              "Name": "Akihira77"
            },
            "Message": "feat: refactore code",
            "Distinct": true,
            "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli/commits/b214e62ad5fe6034f0f9f5614f9fef7bd0a74325"
          }
        ]
      },
      "Repo": {
        "ID": 982508870,
        "Name": "roadmap-sh-Backend/task-tracker-cli",
        "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli"
      },
      "Actor": {
        "ID": 83736692,
        "Login": "Akihira77",
        "AvatarURL": "https://avatars.githubusercontent.com/u/83736692?",
        "URL": "https://api.github.com/users/Akihira77"
      },
      "Org": {
        "ID": 211434693,
        "Login": "roadmap-sh-Backend",
        "AvatarURL": "https://avatars.githubusercontent.com/u/211434693?",
        "URL": "https://api.github.com/orgs/roadmap-sh-Backend"
      },
      "CreatedAt": "2025-05-14T02:05:01Z"
    },
    {
      "ID": 49645238825,
      "Type": "PushEvent",
      "Public": true,
      "Payload": {
        "PushID": 24259529995,
        "Size": 1,
        "DistinctSize": 1,
        "Ref": "refs/heads/master",
        "Head": "4a6e00833bcbc993344189fb5864b3e8494df0e3",
        "Before": "fd14a59a8e559da1dc8575e865fd66855df9d205",
        "Commits": [
          {
            "Sha": "4a6e00833bcbc993344189fb5864b3e8494df0e3",
            "Author": {
              "Email": "tik.dikawahyu123@gmail.com",
              "Name": "Akihira77"
            },
            "Message": "feat: develop list, add, update, delete feature, and setup makefile",
            "Distinct": true,
            "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli/commits/4a6e00833bcbc993344189fb5864b3e8494df0e3"
          }
        ]
      },
      "Repo": {
        "ID": 982508870,
        "Name": "roadmap-sh-Backend/task-tracker-cli",
        "URL": "https://api.github.com/repos/roadmap-sh-Backend/task-tracker-cli"
      },
      "Actor": {
        "ID": 83736692,
        "Login": "Akihira77",
        "AvatarURL": "https://avatars.githubusercontent.com/u/83736692?",
        "URL": "https://api.github.com/users/Akihira77"
      },
      "Org": {
        "ID": 211434693,
        "Login": "roadmap-sh-Backend",
        "AvatarURL": "https://avatars.githubusercontent.com/u/211434693?",
        "URL": "https://api.github.com/orgs/roadmap-sh-Backend"
      },
      "CreatedAt": "2025-05-13T03:59:05Z"
    }
  ]
}
```
