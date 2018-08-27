### /agency/admin/authorization/list

权限列表返回

```json

{
    "Code": 0,
    "Data": {
        "StaffRole": {
            "RoleId": 2,
            "OrganizationId": 1,
            "RoleName": "管理员", //角色名称
            "RoleAuthorization": {
                "1001": {
                    "AuthorizationId": 1001,
                    "Title": "查看成员",
                    "ParentId": 1000,
                    "ParentTitle": "成员管理"
                },
                "1002": {
                    "AuthorizationId": 1002,
                    "Title": "编辑成员",
                    "ParentId": 1000,
                    "ParentTitle": "成员管理"
                }
            },
            "IsDisable": false
        },
        "Modules": [ // 模块权限
            {
                "ModuleAuthorization": { // 模块权限
                    "AuthorizationId": 1000,
                    "Title": "成员管理",  // 权限名称
                    "ParentId": 0,
                    "ParentTitle": "",
                    "IsOwn": false // 是否已勾选
                },
                "SubAuthorizations": { // 子模块权限
                    "1001": {
                        "AuthorizationId": 1001,
                        "Title": "查看成员",
                        "ParentId": 1000,
                        "ParentTitle": "成员管理",
                        "IsOwn": true
                    },
                    "1002": {
                        "AuthorizationId": 1002,
                        "Title": "编辑成员",
                        "ParentId": 1000,
                        "ParentTitle": "成员管理",
                        "IsOwn": true
                    },
                    "1003": {
                        "AuthorizationId": 1003,
                        "Title": "查看角色",
                        "ParentId": 1000,
                        "ParentTitle": "成员管理",
                        "IsOwn": false
                    },
                    "1004": {
                        "AuthorizationId": 1004,
                        "Title": "编辑角色",
                        "ParentId": 1000,
                        "ParentTitle": "成员管理",
                        "IsOwn": false
                    }
                }
            },
            {
                "ModuleAuthorization": {
                    "AuthorizationId": 2000,
                    "Title": "社群管理",
                    "ParentId": 0,
                    "ParentTitle": "",
                    "IsOwn": false
                },
                "SubAuthorizations": {
                    "2001": {
                        "AuthorizationId": 2001,
                        "Title": "查看社群",
                        "ParentId": 2000,
                        "ParentTitle": "社群管理",
                        "IsOwn": false
                    },
                    "2002": {
                        "AuthorizationId": 2002,
                        "Title": "编辑社群",
                        "ParentId": 2000,
                        "ParentTitle": "社群管理",
                        "IsOwn": false
                    }
                }
            },
            {
                "ModuleAuthorization": {
                    "AuthorizationId": 3000,
                    "Title": "团购任务管理",
                    "ParentId": 0,
                    "ParentTitle": "",
                    "IsOwn": false
                },
                "SubAuthorizations": {
                    "2002": {
                        "AuthorizationId": 3002,
                        "Title": "编辑团购任务",
                        "ParentId": 3000,
                        "ParentTitle": "团购任务管理",
                        "IsOwn": false
                    },
                    "3001": {
                        "AuthorizationId": 3001,
                        "Title": "查看团购任务",
                        "ParentId": 3000,
                        "ParentTitle": "团购任务管理",
                        "IsOwn": false
                    }
                }
            }
        ]
    },
    "Message": "success"
}
```

