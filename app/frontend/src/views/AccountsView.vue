<script setup lang="ts">
import {ref} from "vue";
// @ts-ignore
import anime from "animejs";
import {AddAccount, GetAccounts, GetSelectedAccount, RemoveAccount, SelectOneAccount} from "@go/views/AccountsView";
import {closeToast, showDialog, showLoadingToast, showNotify} from 'vant';
import {views} from "@go/models";
import {useRouter} from "vue-router";
import {useI18n} from "vue-i18n";
import LocalStorage from "@/stores/ls";

const showAccounts = ref(true);
const inputToken = ref('');
const isLoading = ref(false);
const accountList = ref(<views.GetAccountsRespAccount[]>[]);
let router = useRouter();

const toShowAddAccount = () => {
  showAccounts.value = false;
  isLoading.value = false;
  anime({
    // 动画
    targets: '.container',
    opacity: [0, 1],
    duration: 800,
    easing: 'easeInOutQuad'
  });
}
const toShowSelectAccount = () => {
  loadAccounts();
  showAccounts.value = true;
  inputToken.value = '';
  anime({
    // 动画
    targets: '.container',
    opacity: [0, 1],
    duration: 800,
    easing: 'easeInOutQuad'
  });
}
const onSubmit = () => {
  console.log('submit');
  isLoading.value = true;
  const loading = showLoadingToast({
    message: t('添加中...'),
    duration: 0,
    forbidClick: true,
    loadingType: 'spinner',
  });
  let req = new views.AddAccountReq;
  req.token = inputToken.value;
  AddAccount(req).then((resp) => {
    if (resp.ok) {
      showNotify({type: 'success', message: t('添加成功')});
      closeToast();
      // 1s后隐藏
      setTimeout(() => {
        toShowSelectAccount();
      }, 1000);
    } else {
      // 1s后隐藏
      isLoading.value = false;
      showNotify({type: 'danger', message: t(resp.error)});
      closeToast();
      setTimeout(() => {
      }, 1000);
    }
  });
}

const loadAccounts = () => {
  // 加载账号
  GetAccounts().then((resp) => {
    if (resp.ok) {
      console.log(resp.accounts);
    } else {
      showNotify({type: 'danger', message: resp.error});
    }
    // mock 很多
    let mockList = <views.GetAccountsRespAccount[]>[];
    resp.accounts.forEach((account) => {
      mockList.push(account);
    })
    // accountList.value = resp.accounts;
    accountList.value = mockList;
  });
}

const onClickAccountInfo = (account: views.GetAccountsRespAccount) => {
  SelectOneAccount(account.id).then((resp) => {
    if (resp.ok) {
      // 动画退出
      anime({
        // 动画
        targets: '.page',
        opacity: [1, 0],
        duration: 260,
        easing: 'easeInOutQuad'
      });
      // 500ms后跳转
      setTimeout(() => {
        router.push('/home');
      }, 300);
    } else {
      showNotify({type: 'danger', message: resp.error});
    }
  });
}

loadAccounts();

const checkSelectedAccount = () => {
  //先检查有没有同意用户协议
  if (LocalStorage.ifAgree()) {
    router.push('/userAgreement');
    return;
  }
  GetSelectedAccount().then((resp) => {
    if (resp.ok) {
      if (resp.id > 0) {
        router.push('/home');
      }
    }
  });
}

checkSelectedAccount();

const showRemoveConfirmDialog = (account: views.GetAccountsRespAccount) => {
  showDialog({
    title: t('提示'),
    message: t('确定要移除吗？'),
    showCancelButton: true,
    confirmButtonText: t('确定'),
    cancelButtonText: t('取消'),
    theme: 'round-button',
  }).then(() => {
    RemoveAccount(account.id).then((resp) => {
      if (resp.ok) {
        showNotify({type: 'success', message: t('移除成功')});
        loadAccounts();
      } else {
        showNotify({type: 'danger', message: resp.error});
      }
    });
  }).catch(() => {
  });
}


const {t} = useI18n();
</script>

<template>
  <div class="page">
    <div class="container" v-show="showAccounts">
      <img src="@/assets/images/appicon.png" alt="" class="logoImg"/>
      <div class="logoAndAccountListContainer">
        <!--        <div class="logoContainer">-->
        <!--          <img src="@/assets/images/appicon.png" alt="" class="logoContainerLogo"/>-->
        <!--        </div>-->

        <div class="AccountListContainer">
          <div class="AccountListContainer2">
            <div
                class="avatarDivAndAccountInfoDiv" v-for="(account, index) in accountList" :key="account.id">
              <!--左边是头像，右边是账号信息-->
              <!--头像，是首字母加背景色组成的，圆形-->
              <div class="avatarDiv">
                <span class="avatarSpan">{{ account.first_name[0] }}</span>
              </div>
              <!--账号信息-->
              <div class="accountInfoDiv"
                   @click="onClickAccountInfo(account)"
              >
                <div class="name">{{ account.first_name }} {{ account.last_name }}</div>
                <div class="runningState"
                     :style="{color: account.running ? '#77d571' : '#999999'}"
                >
                  {{ account.running ? t('运行中') : t('未运行') }}
                </div>
                <div class="splitLine">&nbsp;</div>
              </div>
              <!--跳转图标-->
              <div class="jumpIconDiv" @click="showRemoveConfirmDialog(account)">
                {{ t('移除') }}
              </div>
            </div>

          </div>
          <div class="AccountListContainer3">
            <a class="superLink"
               @click="toShowAddAccount"
            >
              {{ t('添加一个') }}
            </a>
          </div>
        </div>
      </div>
    </div>
    <div class="container" v-show="!showAccounts">
      <img src="@/assets/images/appicon.png" alt="" class="logoImg"/>
      <h2>Robot For Telegram</h2>
      <span class="nbspLine">&nbsp;</span>
      <div class="addAccountTip">
        {{ t('addAccountTip') }}
      </div>
      <div class="inputDiv">
        <input
            type="text" :placeholder="t('请输入Token')"
            v-model="inputToken"
            @keydown.enter="onSubmit"
            :disabled="isLoading"
        />
      </div>
      <span class="nbspLine">&nbsp;</span>
      <span class="nbspLine">&nbsp;</span>
      <span class="nbspLine">&nbsp;</span>
      <a class="superLink"
         @click="toShowSelectAccount"
      >{{ t('选择已有账号') }}</a>
    </div>
  </div>
</template>

<style scoped lang="scss">
//avatarDivAndAccountInfoDiv 中的三个div左中右布局
.avatarDivAndAccountInfoDiv {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  cursor: pointer;
}

//avatarDiv 头像是首字母加背景色组成的，圆形，背景色渐变aabefa - 938ff8
.avatarDiv {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-right: 12px;
  background: linear-gradient(180deg, #aabefa 0%, #938ff8 100%);
}

//accountInfoDiv 账号信息，最上面是First Name + Last Name，下面是id
.accountInfoDiv {
  width: 272px;
  height: 40px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;

  .name {
    font-size: 12px;
    color: #dfdfdf;
    width: 272px;
    height: 12px;
    margin-top: 7px;
  }

  .runningState {
    font-size: 10px;
    width: 272px;
    height: 10px;
    margin-bottom: 10px;
  }

  .splitLine {
    width: 272px;
    height: 1px;
    background: #282828;
  }
}

//jumpIconDiv 跳转图标
.jumpIconDiv {
  width: 40px;
  height: 40px;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  color: #eb4d3d;
  font-size: 12px;
}

// page是整个页面，页面内的container上下左右居中，container中的所有元素左右居中
// page中的文字 默认颜色为#dfdfdf 大小为12px weight为100

.page {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  color: #dfdfdf;
  font-size: 12px;
  font-family: Arial, serif;
  font-weight: normal;
  background: #282828;

  .container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
}

// logoAndAccountListContainer 中的两个div左右布局，宽度为600px
// logoContainer 宽度为300px 高度为300px 所有元素左右居中
// AccountListContainer 宽度为300px 高度自适应 所有元素左右居中
.logoAndAccountListContainer {
  display: flex;
  justify-content: space-between;
  width: 400px;
  margin-bottom: 12px;
  height: 300px;
}

.logoContainer {
  width: 300px;
  height: 300px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.logoContainerLogo {
  width: 300px;
  height: 300px;
}

.logoAndAccountListSplit {
  width: 2px;
  height: 300px;
  background-color: #3d3d3d;
}

// AccountListContainer 包括AccountListContainer2和AccountListContainer3 上下布局
.AccountListContainer {
  width: 400px;
  background-color: #323232;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 0;
  border-radius: 10px;
  overflow: hidden;
  //box-shadow: 0 0 10px 0 #000000;
}

//AccountListContainer2 有固定高度，允许很多列表，上下排列，竖有滚动条，横没有
.AccountListContainer2 {
  width: 400px;
  height: 260px;
  display: flex;
  flex-direction: column;
  //justify-content: center;
  align-items: center;
  padding: 12px;
  overflow-y: auto;
  overflow-x: hidden;
  margin-bottom: 12px;

  div {
    margin-bottom: 12px;
  }
}

.AccountListContainer3 {
  width: 400px;
  display: flex;
  justify-content: center;
  align-items: center;

  a {
    color: #3478f6;
    cursor: pointer;
  }
}

// superLink是暂无账号的提示文字，字体颜色为#dfdfdf
.superLink {
  color: #3478f6;
  cursor: pointer;
}

// logoImg 128px * 128px
.logoImg {
  width: 128px;
  height: 128px;
  margin-bottom: 12px;
}

//h2 大小：20px 颜色：#dfdfdf
h2 {
  font-size: 20px;
  color: #dfdfdf;
}

//nbspLine 20px高
.nbspLine {
  display: block;
  height: 12px;
}

//addAccountTip 230宽 高度自适应
//里面的字 颜色：#9e9e9e 大小：12px
.addAccountTip {
  width: 230px;
  text-align: center;
  margin-bottom: 12px;
  color: #9e9e9e;
  font-size: 12px;
}

//inputDiv 宽280 高40 #323232 圆角10px
.inputDiv {
  width: 280px;
  height: 40px;
  margin-bottom: 12px;

  input {
    width: 100%;
    height: 100%;
    background-color: #323232;
    border: none;
    color: #dfdfdf;
    font-size: 12px;
    padding-left: 10px;
    border-radius: 10px;
  }
}

</style>