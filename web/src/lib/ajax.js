import axios from "axios"
import app from "../config/app"
import token from "../lib/token"
import c from "../lib/const"
import AuthApi from "../api/auth";

class Ajax {
    constructor() {
        let t = token.get(app.auth_token_key)

        this.axios = axios.create({
                        baseURL: app.url_prefix,
                        timeout: 1000,
                        xsrfHeaderName: 'X-CSRFToken',
                        xsrfCookieName: 'csrftoken',
                        withCredentials: true,
                        headers: {
                            Authorization: t,
                            post: {
                                'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
                            }
                        }
                    });
        this.axios.interceptors.response.use((response) => {
            if (response.data.code == c.unauthorized) {
                AuthApi.logout()
            }
            return response
        })
    }

    get(url) {
        return this.axios.get(url)
    }

    post(url, data = {}) {
        return this.axios.post(url, data)
    }

}

export default new Ajax