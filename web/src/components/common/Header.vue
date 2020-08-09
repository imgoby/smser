<template>
    <el-header style="">
        <el-row>
            <el-col :span="17">
                <el-breadcrumb separator="/">
                    <el-breadcrumb-item v-for="(item, index) in breadcrumb" :key="index" :to="{ path: item.getBread().href }">{{item.getBread().name}}</el-breadcrumb-item>
                </el-breadcrumb>
            </el-col>
            <el-col :span="7" class="userinfo">
                <span>{{user.username}}</span>
                <span @click="logout()">退出</span>
            </el-col>
        </el-row>
    </el-header>
</template>

<script>
import { Bread } from "@/lib/BreadcrumbUtil";
import AuthApi from "@/api/auth";
import routes_config from "@/config/routes";

export default {
    data() {
        return {
            user: {username: "demo"}
        }
    },
    props: {
        'breadcrumb': Array
    },
    mounted() {
        this.user = AuthApi.user()
    },
    created() {
        this.initBreadcrumb();
    },
    updated() {
        this.initBreadcrumb();
    },
    methods: {
        initBreadcrumb() {
            this.breadcrumb = this.breadcrumb ? this.breadcrumb : [new Bread("首页", "/")];
        },
        logout() {
            if(AuthApi.logout()) {
                this.$router.push({name: routes_config.LOGIN_ROUTE_NAME});
            }
        }
    }
}
</script>

<style scoped lang="scss">
  @import "@/assets/css/variables.sass"; 
 .el-header {
    text-align: right;
    height: 50px!important;
    font-size: 12px!important;
    line-height: 50px;
    border-bottom: 1px solid rgb(238, 241, 246);
 }
 .el-breadcrumb{
    font-size: 12px!important;
    height: 100%!important;
    line-height: inherit!important;
 }
 .menu-icon{
    padding-top: 5px;
    & i{
        font-size: 25px;margin-right: 10px;
    }
 }
 .userinfo {
    color: $blue-color;
    & span:last-child {
        margin-left: 10px;
        cursor: pointer;
    }
 }
</style>