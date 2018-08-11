import m from "mithril";
import User from "../model/User";
import Signup from "../model/Signup";

var Login = {
    current: {},
    load: function(id) {
        return m
            .request({
                method: "GET",
                url: "/api/profile",
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("token")
                },
                withCredentials: false
            })
            .then(function(result) {
                if (result) {
                    if (result.username) {
                        User.current.username = result.username;
                    } else if (result.error) {
                        User.current.error = result.error;
                        return;
                    }
                    User.current.error = "";
                }
            });
    },
    post: function() {
        return m
            .request({
                method: "POST",
                url: "/api/login",
                data: User.current,
                withCredentials: true
            })
            .then(function(result) {
                if (result && result.token) {
                    localStorage.setItem("token", result.token);
                    //m.route.set("/"); does not load the routes again
                    location.href = "/";
                }
            });
    },
    signup: function() {
        return m
            .request({
                method: "POST",
                url: "/api/signup",
                data: Signup.current,
                withCredentials: true
            })
            .then(function(result) {
                console.log(result.status);
                if (result && result.error) {
                    Signup.current.error = result.error;
                } else if (result && result.status) {
                    localStorage.setItem("token", result.token);
                    Signup.current.error = "";
                    m.route.set("/login");
                }
            });
    }
};

module.exports = Login;
