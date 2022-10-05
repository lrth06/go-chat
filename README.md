# Go-Chat

<pre>
├── client
│ ├── build
│ │ └── ...
│ ├── package.json
│ ├── package-lock.json
│ ├── postcss.config.js
│ ├── public
│ │ ├── images
│ │ │ ├── 404.svg
│ │ │ ├── chat.svg
│ │ │ ├── icon.svg
│ │ │ └── support-team.svg
│ │ ├── index.html
│ │ ├── manifest.json
│ │ └── robots.txt
│ ├── src
│ │ ├── App.jsx
│ │ ├── App.test.jsx
│ │ ├── components
│ │ │ ├── alert
│ │ │ │ ├── error.jsx
│ │ │ │ ├── info.jsx
│ │ │ │ └── success.jsx
│ │ │ ├── chat
│ │ │ │ ├── Chat.jsx
│ │ │ │ ├── MessageBar.jsx
│ │ │ │ ├── Message.jsx
│ │ │ │ ├── UserDropdown.jsx
│ │ │ │ └── UserList.jsx
│ │ │ ├── Layout.jsx
│ │ │ └── TimeAgo.jsx
│ │ ├── context
│ │ │ ├── ThemeContext.jsx
│ │ │ └── UserContext.jsx
│ │ ├── css
│ │ │ └── styles.css
│ │ ├── hooks
│ │ │ ├── useForm.jsx
│ │ │ └── useWebsocket.jsx
│ │ ├── index.jsx
│ │ ├── setupTests.jsx
│ │ ├── utils
│ │ │ ├── goToRoom.jsx
│ │ │ ├── logout.jsx
│ │ │ └── parseJwt.jsx
│ │ └── views
│ │ ├── auth
│ │ │ ├── Login.jsx
│ │ │ ├── PasswordReset.jsx
│ │ │ └── Register.jsx
│ │ ├── blog
│ │ │ ├── EditPost.jsx
│ │ │ ├── PostDirectory.jsx
│ │ │ └── Post.jsx
│ │ ├── error
│ │ │ ├── Forbidden.jsx
│ │ │ ├── NotFound.jsx
│ │ │ ├── ServerError.jsx
│ │ │ └── Unauthorized.jsx
│ │ ├── Home.jsx
│ │ ├── pricing
│ │ │ └── Base.jsx
│ │ ├── rooms
│ │ │ ├── Directory.jsx
│ │ │ ├── Editor.jsx
│ │ │ └── Room.jsx
│ │ └── user
│ │ ├── EditProfile.jsx
│ │ ├── Profile.jsx
│ │ ├── PublicProfile.jsx
│ │ └── Recover.jsx
│ ├── tailwind.config.js
│ └── yarn.lock
├── go.mod
├── go.sum
├── lib
│ ├── handlers
│ │ ├── api_handlers.go
│ │ ├── auth
│ │ │ ├── login.go
│ │ │ └── logout.go
│ │ ├── room_handlers.go
│ │ ├── users
│ │ │ ├── create_user.go
│ │ │ ├── delete_user.go
│ │ │ ├── get_user.go
│ │ │ ├── get_users.go
│ │ │ └── update_user.go
│ │ └── websocket_handlers.go
│ ├── middleware
│ │ ├── auth.go
│ │ ├── logs.go
│ │ ├── self_or_admin.go
│ │ ├── token.go
│ │ └── validation
│ │ ├── room_validation.go
│ │ └── user_validation.go
│ ├── models
│ │ ├── post_model.go
│ │ ├── room_model.go
│ │ └── user_model.go
│ ├── routes
│ │ └── routes.go
│ ├── structs
│ │ ├── structs.go
│ │ └── websocket.go
│ └── utils
│ ├── config
│ │ ├── config.go
│ │ └── db.go
│ ├── logger.go
│ ├── shutdown.go
│ └── startup.go
├── LICENSE
├── main.go
├── main_test.go
├── makefile
├── README.md
└── sample.env
</pre>
