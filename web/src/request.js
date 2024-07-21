import axios from "axios";
import {ElNotification} from "element-plus";
import router from "@/router/index.js";

const request = axios.create({
    baseURL: "/api",
})

request.interceptors.request.use(
    config => {
        if (config.method === "post" || config.method === "put") {
            config.headers.set("Content-Type", "multipart/form-data")
        }

        return config
    },
    error => {
        console.log(error)
        return error
    }
)

request.interceptors.response.use(
    response => {
        return response.data
    },
    error => {
        if (error.response) {
            switch (error.response.status) {
                case 400:
                case 422:
                    ElNotification({
                        title: 'Error',
                        message: error.response.data.message,
                        type: 'error',
                    })
                    break
                case 401:
                    localStorage.removeItem("isAuth");
                    router.push("/login").then(r => {console.log("跳转 /login")})
            }
        } else {
            ElNotification({
                title: 'Error',
                message: "请求错误",
                type: 'error',
            })
        }

        return error
    }
)

export default request
