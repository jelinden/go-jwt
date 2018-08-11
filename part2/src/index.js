import m from "mithril";
import Login from "./view/Login";
import Signup from "./view/Signup";
import Protected from "./view/Protected";
import NotProtected from "./view/NotProtected";
import Layout from "./view/Layout";
import Profile from "./service/Login";
import User from "./model/User";

import "./index.css";

Profile.load().then(function() {
    if (User.current.username) {
        AppRouter();
    } else {
        LoginRouter();
    }
});

function AppRouter() {
    m.route(document.body, "/", {
        "/": {
            render: function() {
                return m(Layout, m(Protected));
            }
        },
        "/login": {
            render: function() {
                return m(Layout, m(Login));
            }
        },
        "/signup": {
            render: function() {
                return m(Layout, m(Signup));
            }
        }
    });
}

function LoginRouter() {
    m.route(document.body, "/", {
        "/": {
            render: function() {
                return m(Layout, m(NotProtected));
            }
        },
        "/login": {
            render: function() {
                return m(Layout, m(Login));
            }
        },
        "/signup": {
            render: function() {
                return m(Layout, m(Signup));
            }
        }
    });
}

m.route.prefix("");
