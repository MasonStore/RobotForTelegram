<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {v4 as uuidv4} from 'uuid';
// @ts-ignore
import anime from "animejs";
import {BrowserOpenURL} from "../../wailsjs/runtime";
import {useRouter} from "vue-router";
import {SelectOneAccount} from "@go/views/AccountsView";
import {ref} from "vue";
import {closeToast, showLoadingToast, showNotify, showToast} from 'vant';
import {GetHome, SaveHome, SetRunning} from "@go/views/HomeView";
import {views} from "@go/models";

interface KeyboardGroup {
  id: string;
  name: string;
  keys: string[][];
}

interface Button {
  Title: string;
  Url: string;
}

interface Game {
  GameShortName: string;
  GameTitle: string;
  Url: string;
  ButtonText: string;
}

interface Pic {
  Base64: string;
}

interface Text {
  Content: string;
}

interface Reply {
  id: string;
  name: string;
  isTemplate: string;
  messageType: string; //text, pic, game
  //game可能为null
  gameContent: Game;
  picContent: Pic;
  textContent: Text;
  expandType: string; //none, keyboardGroup, button
  keyboardGroupId: string;
  keyboardGroupName: string | undefined | null;
  buttons: Button[];
}

interface Command {
  id: string;
  command: string;
  replyId: string;
  replyName: string | null | undefined;
  replyIsTemplate: string | null | undefined;
  hookUrl: string | null | undefined;
}

interface AdvancedConfig {
  ChatGptEnabled: boolean;
  ChatGptApiKey: string;
  ChatGptHttpProxy: string;
  ChatGptModel: string;
  ChatGptSystemPrompt: string;
  ChatGptBaseUrl: string;
  ChatGptTemperature: string;
  ChatGptMaxTokens: string;
  ChatGptTopP: string;
  ChatGptPresencePenalty: string;
  ChatGptFrequencyPenalty: string;
  ChatGptHistoryCount: string;
  ChatGptTimeout: string;
}

interface RobotInfo {
  FirstName: string;
  LastName: string;
  Username: string;
  Token: string;
}

const advancedConfig = ref<AdvancedConfig>({
  ChatGptEnabled: false,
  ChatGptApiKey: "",
  ChatGptHttpProxy: "",
  ChatGptModel: "gpt-3.5-turbo",
  ChatGptSystemPrompt: "You are a friendly AI created by OpenAI. You are very helpful and will answer any questions you ask.",
  ChatGptBaseUrl: "",
  ChatGptTemperature: "0.5",
  ChatGptMaxTokens: "150",
  ChatGptTopP: "1.0",
  ChatGptPresencePenalty: "0",
  ChatGptFrequencyPenalty: "0",
  ChatGptHistoryCount: "4",
  ChatGptTimeout: "60",
});
const keyboardGroups = ref(<KeyboardGroup[]>[]);
const commands = ref(<Command[]>[]);
const replies = ref(<Reply[]>[]);
const running = ref(false);
const robotInfo = ref<RobotInfo>({
  FirstName: "",
  LastName: "",
  Username: "",
  Token: "",
})

const getHome = () => {
  GetHome().then((resp) => {
    if (resp.ok) {
      const data = resp.data;
      running.value = data.running;
      robotInfo.value = {
        FirstName: data.robotInfo.firstName,
        LastName: data.robotInfo.lastName,
        Username: data.robotInfo.username,
        Token: data.robotInfo.token,
      };
      advancedConfig.value = {
        ChatGptEnabled: data.advancedConfig.ChatGptEnabled,
        ChatGptApiKey: data.advancedConfig.ChatGptApiKey,
        ChatGptHttpProxy: data.advancedConfig.ChatGptHttpProxy,
        ChatGptModel: data.advancedConfig.ChatGptModel,
        ChatGptSystemPrompt: data.advancedConfig.ChatGptSystemPrompt,
        ChatGptBaseUrl: data.advancedConfig.ChatGptBaseUrl,
        ChatGptTemperature: data.advancedConfig.ChatGptTemperature,
        ChatGptMaxTokens: data.advancedConfig.ChatGptMaxTokens,
        ChatGptTopP: data.advancedConfig.ChatGptTopP,
        ChatGptPresencePenalty: data.advancedConfig.ChatGptPresencePenalty,
        ChatGptFrequencyPenalty: data.advancedConfig.ChatGptFrequencyPenalty,
        ChatGptHistoryCount: data.advancedConfig.ChatGptHistoryCount,
        ChatGptTimeout: data.advancedConfig.ChatGptTimeout,
      };
      let keyboardGroupsTmp = <KeyboardGroup[]>[];
      data.keyboardGroups.forEach((item: any) => {
        keyboardGroupsTmp.push({
          id: item.id,
          name: item.name,
          keys: item.keys,
        });
      });
      keyboardGroups.value = keyboardGroupsTmp;
      let commandsTmp = <Command[]>[];
      data.commands.forEach((item: any) => {
        commandsTmp.push({
          id: item.id,
          command: item.command,
          replyId: item.replyId,
          replyName: item.replyName,
          replyIsTemplate: item.replyIsTemplate,
          hookUrl: item.hookUrl,
        });
      });
      commands.value = commandsTmp;
      let repliesTmp = <Reply[]>[];
      data.replies.forEach((item: any) => {
        repliesTmp.push({
          id: item.id,
          name: item.name,
          isTemplate: item.isTemplate,
          messageType: item.messageType,
          gameContent: item.gameContent,
          picContent: item.picContent,
          textContent: item.textContent,
          expandType: item.expandType,
          keyboardGroupId: item.keyboardGroupId,
          keyboardGroupName: item.keyboardGroupName,
          buttons: item.buttons,
        });
      });
      replies.value = repliesTmp;
    } else {
      showNotify({type: 'danger', message: t('获取数据失败')});
    }
  });
}
getHome()
const onSubmitForm = () => {
  //表单校验
  formRef.value.validate().then(() => {
    //校验通过
    const loading = showLoadingToast({
      message: t('保存中...'),
      duration: 0,
      forbidClick: true,
      loadingType: 'spinner',
    });
    let data = new views.GetHomeRespData();
    data.running = running.value;
    data.robotInfo = {
      FirstName: robotInfo.value.FirstName,
      LastName: robotInfo.value.LastName,
      Username: robotInfo.value.Username,
      Token: robotInfo.value.Token,
    };
    data.advancedConfig = {
      ChatGptEnabled: advancedConfig.value.ChatGptEnabled,
      ChatGptApiKey: advancedConfig.value.ChatGptApiKey,
      ChatGptHttpProxy: advancedConfig.value.ChatGptHttpProxy,
      ChatGptModel: advancedConfig.value.ChatGptModel,
      ChatGptSystemPrompt: advancedConfig.value.ChatGptSystemPrompt,
      ChatGptBaseUrl: advancedConfig.value.ChatGptBaseUrl,
      ChatGptTemperature: advancedConfig.value.ChatGptTemperature,
      ChatGptMaxTokens: advancedConfig.value.ChatGptMaxTokens,
      ChatGptTopP: advancedConfig.value.ChatGptTopP,
      ChatGptPresencePenalty: advancedConfig.value.ChatGptPresencePenalty,
      ChatGptFrequencyPenalty: advancedConfig.value.ChatGptFrequencyPenalty,
      ChatGptHistoryCount: advancedConfig.value.ChatGptHistoryCount,
      ChatGptTimeout: advancedConfig.value.ChatGptTimeout,
    };
    data.keyboardGroups = [];
    keyboardGroups.value.forEach((item) => {
      let keyboardGroup = new views.GetHomeRespKeyboardGroup();
      keyboardGroup.id = item.id;
      keyboardGroup.name = item.name;
      keyboardGroup.keys = item.keys;
      data.keyboardGroups.push(keyboardGroup);
    });
    data.commands = [];
    commands.value.forEach((item) => {
      let command = new views.GetHomeRespCommand();
      command.id = item.id;
      command.command = item.command;
      command.replyId = item.replyId;
      command.replyName = item.replyName;
      command.replyIsTemplate = item.replyIsTemplate;
      command.hookUrl = item.hookUrl;
      data.commands.push(command);
    });
    data.replies = [];
    replies.value.forEach((item) => {
      let reply = new views.GetHomeRespReply();
      reply.id = item.id;
      reply.name = item.name;
      reply.isTemplate = item.isTemplate;
      reply.messageType = item.messageType;
      reply.gameContent = item.gameContent;
      reply.picContent = item.picContent;
      reply.textContent = item.textContent;
      reply.expandType = item.expandType;
      reply.keyboardGroupId = item.keyboardGroupId;
      reply.keyboardGroupName = item.keyboardGroupName;
      reply.buttons = item.buttons;
      data.replies.push(reply);
    });
    SaveHome(data).then((resp) => {
      closeToast();
      if (resp.ok) {
        showNotify({type: 'success', message: t('保存成功')});
        getHome();
      } else {
        showNotify({type: 'danger', message: t('保存失败')});
      }
    });
  }).catch(() => {
    //校验不通过
    showNotify({type: 'danger', message: t('请检查参数')});
  });
}

const onRunningChanged = (value: boolean) => {
  SetRunning(value).then((resp) => {
    if (resp.ok) {
      running.value = value;
    } else {
      showNotify({type: 'danger', message: resp.error});
    }
  });
}

const formRef: any = ref<InstanceType<typeof import('vant').Form>>();

const {t} = useI18n();

const onClickLink = () => {
  //https://wails.io/docs/gettingstarted/installation
  BrowserOpenURL("https://wails.io/docs/gettingstarted/installation");
};

const router = useRouter();

const onClickHome = () => {
  SelectOneAccount(0).then((resp) => {
    if (resp.ok) {
      // 动画退出
      anime({
        // 动画
        targets: '.page',
        opacity: [1, 0],
        duration: 260,
        easing: 'easeInOutQuad'
      });
      // 400ms后跳转
      setTimeout(() => {
        router.push('/');
      }, 300);
    } else {
    }
  });
};

const activeNameForKeyboardCollapse = ref('');
const activeNameForReplyCollapse = ref('');
const activeNameForCommandCollapse = ref('');

const copyText = (text: string) => {
  //@ts-ignore
  CommonJS.copyText(text);
  showToast(t('复制成功'));
};

const newKeyboardRow = (index: number) => {
  keyboardGroups.value[index].keys.push(['']);
}

const removeKeyboardRow = (index: number, rowIndex: number) => {
  keyboardGroups.value[index].keys.splice(rowIndex, 1);
}

const newKeyboardCol = (index: number, rowIndex: number) => {
  //一行10个是极限
  if (keyboardGroups.value[index].keys[rowIndex].length >= 10) {
    return;
  }
  keyboardGroups.value[index].keys[rowIndex].push('');
}

const removeKeyboardCol = (index: number, rowIndex: number, keyIndex: number) => {
  keyboardGroups.value[index].keys[rowIndex].splice(keyIndex, 1);
}

const showEditKeyboardGroupNameDialog = ref(false);
const EditKeyboardGroupIndex = ref(0);

const onClickEditKeyboardGroupName = (index: number) => {
  EditKeyboardGroupIndex.value = index;
  showEditKeyboardGroupNameDialog.value = true;
}

const onClickNewKeyboard = () => {
  keyboardGroups.value.push({
    id: uuidv4(),
    name: t("未命名"),
    keys: [
      ['', '', ''],
      ['', '', ''],
      ['', '', ''],
    ]
  });
}

const onClickNewReply = () => {
  replies.value.push({
    id: uuidv4(),
    name: t("未命名"),
    isTemplate: null,
    messageType: "text",
    gameContent: {
      GameShortName: "",
      GameTitle: "",
      Url: "",
      ButtonText: "",
    },
    picContent: {
      Base64: ""
    },
    textContent: {Content: ""},
    expandType: "none",
    keyboardGroupId: "",
    buttons: <Button[]>[],
  });
}

const onClickUploadPicButton = (index: number) => {
  // @ts-ignore
  const el = document.getElementById('inputPic' + index.toString());
  el.click();
  el.onchange = (e: any) => {
    const file = e.target.files[0];
    // 判断是否是图片 png gif jpg jpeg
    if (!/image\/\w+/.test(file.type)) {
      showNotify({type: 'danger', message: t('请上传图片')});
      return false;
    }
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = (e: any) => {
      replies.value[index].picContent.Base64 = e.target.result;
    };
  };
}

const showKeyboardGroupPicker = ref(false);
const showKeyboardGroupPickerIndex = ref(0);
const keyboardGroupColumns = ref([
  {text: '', value: ''},
]);

const onClickShowKeyboardGroupPicker = (index: number) => {
  showKeyboardGroupPicker.value = true;
  showKeyboardGroupPickerIndex.value = index;
  let keyboardGroupColumnsTmp = [{
    text: t('无'),
    value: '',
  }];
  keyboardGroups.value.forEach((item, index) => {
    keyboardGroupColumnsTmp.push({text: item.name, value: item.id.toString()});
  });
  keyboardGroupColumns.value = keyboardGroupColumnsTmp;
}

const onConfirmKeyboardGroupPicker = ({selectedOptions}) => {
  showKeyboardGroupPicker.value = false;
  replies.value[showKeyboardGroupPickerIndex.value].keyboardGroupId = selectedOptions[0].value;
  replies.value[showKeyboardGroupPickerIndex.value].keyboardGroupName = selectedOptions[0].text;
}

const onClickNewCommand = () => {
  commands.value.push({
    id: uuidv4(),
    command: t("未命名"),
    replyId: 0,
  });
}
const showReplyPicker = ref(false);
const showReplyPickerIndex = ref(0);
const replyColumns = ref([
  {text: '', value: ''},
]);

const onClickShowReplyPicker = (index: number) => {
  showReplyPicker.value = true;
  showReplyPickerIndex.value = index;
  let replyColumnsTmp = [];
  replies.value.forEach((item, index) => {
    replyColumnsTmp.push({text: item.name, value: item.id.toString()});
  });
  replyColumns.value = replyColumnsTmp;
}

const onConfirmReplyPicker = ({selectedOptions}) => {
  showReplyPicker.value = false;
  commands.value[showReplyPickerIndex.value].replyId = selectedOptions[0].value;
  commands.value[showReplyPickerIndex.value].replyName = selectedOptions[0].text;
  let selectedReply = replies.value.find((item) => item.id === selectedOptions[0].value);
  if (selectedReply) {
    commands.value[showReplyPickerIndex.value].replyIsTemplate = selectedReply.isTemplate;
  }
}

const inputValidatorMaxToken = (value) => {
  if (parseInt(value) < 1 || parseInt(value) > 2048) {
    return Promise.resolve(false)
  }
  return Promise.resolve(true);
}

const inputValidatorHistoryCount = (value) => {
  if (parseInt(value) < 1 || parseInt(value) > 20) {
    return Promise.resolve(false)
  }
  return Promise.resolve(true);
}
const inputValidatorTimeout = (value) => {
  if (parseInt(value) < 1 || parseInt(value) > 600) {
    return Promise.resolve(false)
  }
  return Promise.resolve(true);
}

const inputValidatorTemperature = (value) => {
  // 0-1
  if (parseFloat(value) < 0 || parseFloat(value) > 1) {
    return Promise.resolve(false)
  }
  return Promise.resolve(true);
}

const chatGPTModels = ref([
  'gpt-4',
  'gpt-3.5-turbo',
  "gpt-4-32k-0613",
  "gpt-4-32k-0314",
  "gpt-4-32k",
  "gpt-4-0613",
  "gpt-4-0314",
  "gpt-4-0125-preview",
  "gpt-4-1106-preview",
  "gpt-4-turbo-preview",
  "gpt-4-vision-preview",
  "gpt-3.5-turbo-0125",
  "gpt-3.5-turbo-1106",
  "gpt-3.5-turbo-0613",
  "gpt-3.5-turbo-0301",
  "gpt-3.5-turbo-16k",
  "gpt-3.5-turbo-16k-0613",
  "gpt-3.5-turbo-instruct",
])
const chatGPTModelColumns = ref([]);
chatGPTModelColumns.value = chatGPTModels.value.map((item) => {
  return {
    text: item,
    value: item,
  }
})
const showSelectChatGPTModels = ref(false);
const onConfirmSelectChatGPTModel = ({selectedOptions}) => {
  showSelectChatGPTModels.value = false;
  advancedConfig.value.ChatGptModel = selectedOptions[0].value;
}

const onClickService = () => {
  BrowserOpenURL("https://t.me/RobotForTGSupport")
}
</script>

<template>
  <van-nav-bar
      :title="t('设置')"
      :left-text="t('主页')"
      left-arrow
      @click-left="onClickHome"
  />
  <van-dialog v-model:show="showEditKeyboardGroupNameDialog"
              :show-confirm-button="false"
              close-on-click-overlay
  >
    <div class="editKeyboardGroupNameDialogDiv">
      <h2 style="color: #dfdfdf">{{ t('修改键盘组名称') }}</h2>
      <input v-model="keyboardGroups[EditKeyboardGroupIndex].name" :placeholder="t('输入键盘组名称')"
             class="editKeyboardGroupNameInput"/>
    </div>
  </van-dialog>
  <div class="page">
    <van-form class="form" @submit="onSubmitForm" required="auto"
              validate-trigger="onSubmit"
              ref="formRef"
    >
      <div class="container">
        <h2 class="van-doc-demo-block__title">{{ t('基本信息') }}</h2>
        <van-cell-group>
          <van-field
              :label-width="50"
              name="switch" :label="t('启动')">
            <template #input>
              <van-switch
                  @update:model-value="onRunningChanged"
                  active-color="#6bd35f" :model-value="running" size="16px"/>
            </template>
          </van-field>
          <van-cell title="FirstName" :value="robotInfo.FirstName"/>
          <van-cell title="LastName" :value="robotInfo.LastName"/>
          <van-cell
              @click="copyText(robotInfo.Username)" class="cursor-pointer" :title="t('用户名')"
              :value="robotInfo.Username"/>
          <van-cell
              @click="copyText(robotInfo.Token)"
              class="cursor-pointer" title="Token" :label="robotInfo.Token"/>
        </van-cell-group>

        <h2 class="van-doc-demo-block__title">
          {{ t('键盘设置') }}
          <van-icon name="plus" class="cursor-pointer" @click="onClickNewKeyboard"/>
        </h2>
        <van-collapse v-model="activeNameForKeyboardCollapse" accordion>
          <van-collapse-item
              v-for="(item, index) in keyboardGroups" :key="index" :name="index.toString()"
          >
            <template #title>
              <div>
                {{ item.name }}&nbsp;
                <van-icon name="revoke" class="cursor-pointer"
                          color="#eb4d3d"
                          @click="keyboardGroups.splice(index, 1)"
                />
              </div>
            </template>
            <van-row :gutter="1" justify="center"
                     v-for="(row, rowIndex) in item.keys" :key="rowIndex"
            >
              <van-col class="van-col-keyboard"
                       v-for="(key, keyIndex) in row" :key="keyIndex"
                       :span="parseInt((20/row.length).toString()).toString()"
              >
                <div class="inputAndClose">
                  <van-icon name="clear" @click="removeKeyboardCol(index, rowIndex, keyIndex)"/>
                  <input
                      :placeholder="t('输入键盘文字')"
                      v-model="keyboardGroups[index].keys[rowIndex][keyIndex]" class="keyInput"/>
                </div>
              </van-col>
              <van-col class="van-col-keyboard" span="2">
                <a class="superLink"
                   @click="newKeyboardCol(index, rowIndex)"
                >{{ t('新增列') }}</a>
              </van-col>
              <van-col class="van-col-keyboard" span="2">
                <a class="superLinkDanger"
                   @click="removeKeyboardRow(index, rowIndex)"
                >{{ t('删除行') }}</a>
              </van-col>
            </van-row>
            <van-divider/>
            <van-row :gutter="1" justify="center">
              <van-col class="van-col-keyboard" span="12">
                <a class="superLink"
                   @click="newKeyboardRow(index)"
                >{{ t('新增行') }}</a>
              </van-col>
              <van-col class="van-col-keyboard" span="12">
                <a class="superLinkDanger"
                   @click="onClickEditKeyboardGroupName(index)"
                >{{ t('修改名称') }}</a>
              </van-col>
            </van-row>
          </van-collapse-item>
        </van-collapse>

        <h2 class="van-doc-demo-block__title">{{ t('回复内容') }}
          <van-icon name="plus" class="cursor-pointer" @click="onClickNewReply"/>
        </h2>
        <van-collapse accordion v-model="activeNameForReplyCollapse">
          <van-collapse-item v-for="(item, index) in replies" :key="index"
                             :name="index.toString()">
            <template #title>
              <div>
                {{ item.name }}&nbsp;
                <van-icon name="revoke" class="cursor-pointer"
                          color="#eb4d3d"
                          @click="replies.splice(index, 1)"
                />
              </div>
            </template>
            <van-cell-group inset>
              <van-field
                  :label-width="200"
                  v-model="replies[index].name"
                  name="别名"
                  :label="t('别名')"
                  :placeholder="t('请输入别名')"
                  :rules="[{ required: true, message: t('请输入别名') }]"
              />
              <van-field
                  :label-width="200"
                  v-model="replies[index].isTemplate"
                  name="是否模板"
                  :label="t('是否模板')"
                  :rules="[{ required: true, message: t('请选择是否是模板') }]"
              >
                <template #input>
                  <van-radio-group direction="horizontal" v-model="replies[index].isTemplate">
                    <van-radio name="true">{{ t('是') }}</van-radio>
                    <van-radio name="false">{{ t('否') }}</van-radio>
                  </van-radio-group>
                </template>
              </van-field>
              <van-field
                  :label-width="200"
                  v-model="replies[index].messageType"
                  name="消息类型"
                  :label="t('消息类型')"
                  :rules="[{ required: true, message: t('请选择消息类型') }]"
              >
                <template #input>
                  <van-radio-group direction="horizontal" v-model="replies[index].messageType">
                    <van-radio name="text">{{ t('文本') }}</van-radio>
                    <van-radio name="pic">{{ t('图片') }}</van-radio>
                    <van-radio name="game">{{ t('游戏') }}</van-radio>
                  </van-radio-group>
                </template>
              </van-field>
              <van-field
                  :label-width="200"
                  v-if="replies[index].isTemplate === 'false' && replies[index].messageType === 'text'"
                  v-model="replies[index].textContent.Content"
                  rows="1"
                  autosize
                  :label="t('文本内容')"
                  type="textarea"
                  :placeholder="t('请输入文本内容')"
                  :rules="[{ required: true, message: t('请输入文本内容') }]"
              />
              <van-field
                  :label-width="200"
                  v-if="replies[index].isTemplate === 'false' && replies[index].messageType === 'game'"
                  v-model="replies[index].gameContent.GameShortName"
                  :label="t('游戏短名')"
                  :placeholder="t('请输入游戏短名')"
                  :rules="[{ required: true, message: t('请输入游戏短名') }]"
              />
              <van-field
                  :label-width="200"
                  v-if="replies[index].isTemplate === 'false' && replies[index].messageType === 'game'"
                  v-model="replies[index].gameContent.GameTitle"
                  :label="t('游戏标题')"
                  :placeholder="t('请输入游戏标题')"
                  :rules="[{ required: true, message: t('请输入游戏标题') }]"
              />
              <van-field
                  :label-width="200"
                  v-if="replies[index].isTemplate === 'false' && replies[index].messageType === 'game'"
                  v-model="replies[index].gameContent.Url"
                  :label="t('游戏链接')"
                  :placeholder="t('请输入游戏链接')"
                  :rules="[{ required: true, message: t('请输入游戏链接') }]"
              />
              <van-field
                  :label-width="200"
                  v-if="replies[index].isTemplate === 'false' && replies[index].messageType === 'game'"
                  v-model="replies[index].gameContent.ButtonText"
                  :label="t('游戏按钮')"
                  :placeholder="t('请输入游戏按钮')"
                  :rules="[{ required: true, message: t('请输入游戏按钮') }]"
              />
              <van-field
                  :label-width="200"
                  v-if="replies[index].isTemplate === 'false' && replies[index].messageType === 'pic'"
                  v-model="replies[index].picContent.Base64"
                  :label="t('图片')"
                  :rules="[{ required: true, message: t('请上传图片') }]"
              >
                <template #input>
                  <img
                      height="100"
                      width="100"
                      :src="replies[index].picContent.Base64"
                      alt=""
                  />
                  <input type="file"
                         accept="image/gif,image/jpeg,image/png,image/jpg"
                         style="display:none;"
                         :id="'inputPic'+index.toString()"/>
                  <van-button color="#76d571" icon="plus"
                              size="mini"
                              plain
                              type="primary"
                              @click="onClickUploadPicButton(index)"
                  >{{ t('上传图片') }}
                  </van-button>
                </template>
              </van-field>
              <van-field :label-width="200"
                         name="扩展类型"
                         :label="t('扩展类型')"
                         v-model="replies[index].expandType"
                         :rules="[{ required: true, message: t('请选择扩展类型') }]"
              >
                <template #input>
                  <van-radio-group direction="horizontal" v-model="replies[index].expandType">
                    <van-radio name="none">{{ t('无') }}</van-radio>
                    <van-radio name="keyboardGroup">{{ t('键盘组') }}</van-radio>
                    <van-radio name="button">{{ t('按钮') }}</van-radio>
                  </van-radio-group>
                </template>
              </van-field>
              <van-field :label-width="200"
                         v-if="replies[index].expandType === 'keyboardGroup'"
                         :label="t('扩展键盘')"
                         :rules="[{ required: true, message: t('请选择扩展键盘') }]"
                         is-link
                         readonly
                         :placeholder="t('选择键盘组')"
                         @click="onClickShowKeyboardGroupPicker(index)"
                         v-model="replies[index].keyboardGroupName"
              />
              <van-popup v-model:show="showKeyboardGroupPicker" round position="bottom">
                <van-picker
                    :columns="keyboardGroupColumns"
                    @cancel="showKeyboardGroupPicker = false"
                    @confirm="onConfirmKeyboardGroupPicker"
                />
              </van-popup>
              <van-field :label-width="200"
                         v-if="replies[index].expandType === 'button'"
                         :placeholder="t('请设置扩展按钮')"
              >
                <template #label>
                  <div>
                    {{ t('扩展按钮') }}&nbsp;
                    <van-icon name="plus" class="cursor-pointer green"
                              @click="replies[index].buttons.push({Title: '', Url: ''})"
                    />
                  </div>
                </template>
                <template #input>
                  <div>
                    <van-row v-for="(button, buttonIndex) in replies[index].buttons" :key="buttonIndex"
                             :gutter="1" justify="center"
                    >
                      <van-col span="10">
                        <van-field :label-width="50"
                                   v-model="replies[index].buttons[buttonIndex].Title"
                                   :label="t('文字')"
                                   :placeholder="t('请输入按钮文字')"
                                   :rules="[{ required: true, message: t('请输入按钮文字') }]"
                        />
                      </van-col>
                      <van-col span="10">
                        <van-field :label-width="50"
                                   v-model="replies[index].buttons[buttonIndex].Url"
                                   :label="t('链接')"
                                   :placeholder="t('请输入按钮链接')"
                                   :rules="[{ required: true, message: t('请输入按钮链接'), pattern: /^http.+/ }]"
                        />
                      </van-col>
                      <van-col span="4">
                        <van-icon name="clear" @click="replies[index].buttons.splice(buttonIndex, 1)"/>
                      </van-col>
                    </van-row>
                  </div>
                </template>
              </van-field>
            </van-cell-group>
          </van-collapse-item>
        </van-collapse>

        <h2 class="van-doc-demo-block__title">{{ t('命令处理') }}
          <van-icon name="plus" class="cursor-pointer" @click="onClickNewCommand"/>
        </h2>
        <van-collapse accordion v-model="activeNameForCommandCollapse">
          <van-collapse-item v-for="(item, index) in commands" :key="index"
                             :name="index.toString()">
            <template #title>
              <div>
                {{ item.command }}&nbsp;
                <van-icon name="revoke" class="cursor-pointer"
                          color="#eb4d3d"
                          @click="commands.splice(index, 1)"
                />
              </div>
            </template>
            <van-cell-group inset>
              <van-field :label-width="200"
                         v-model="commands[index].command"
                         name="命令"
                         :label="t('命令')"
                         :placeholder="t('请输入命令(支持正则表达式)')"
                         :rules="[{ required: true, message: t('请输入命令(支持正则表达式)') }]"
              />
              <van-field :label-width="200"
                         :label="t('回复')"
                         :rules="[{ required: true, message: t('请选择回复内容') }]"
                         is-link
                         readonly
                         :placeholder="t('请选择回复内容')"
                         @click="onClickShowReplyPicker(index)"
                         v-model="commands[index].replyName"
              />
              <van-popup v-model:show="showReplyPicker" round position="bottom">
                <van-picker
                    :columns="replyColumns"
                    @cancel="showReplyPicker = false"
                    @confirm="onConfirmReplyPicker"
                />
              </van-popup>
              <van-field :label-width="200"
                         v-if="commands[index].replyIsTemplate=='true'"
                         v-model="commands[index].hookUrl"
                         name="Hook"
                         label="Hook"
                         :placeholder="t('请输入Hook地址')"
                         :rules="[{ required: true, message: t('请输入Hook地址') }]"
              />
            </van-cell-group>
          </van-collapse-item>

        </van-collapse>

        <h2 class="van-doc-demo-block__title">ChatGPT</h2>
        <van-cell-group>
          <van-field :label-width="200" name="switch" :label="t('启用')">
            <template #input>
              <van-switch
                  active-color="#6bd35f"
                  v-model="advancedConfig.ChatGptEnabled"
                  size="16px"/>
            </template>
          </van-field>
          <van-field :label-width="200"
                     v-if="advancedConfig.ChatGptEnabled"
                     v-model="advancedConfig.ChatGptApiKey"
                     name="ChatGptApiKey"
                     label="api-key"
                     :placeholder="t('请输入api-key')"
                     :rules="[{ required: true, message: t('请输入api-key') }]"
          />
          <van-field :label-width="200"
                     v-if="advancedConfig.ChatGptEnabled"
                     v-model="advancedConfig.ChatGptModel"
                     name="ChatGptModel"
                     :label="t('模型')"
                     :placeholder="t('请选择模型')"
                     :rules="[{ required: true, message: t('请选择模型')}]"
                     @click="showSelectChatGPTModels = true"
          />
          <van-popup v-model:show="showSelectChatGPTModels" position="bottom">
            <van-picker
                :columns="chatGPTModelColumns"
                @confirm="onConfirmSelectChatGPTModel"
                @cancel="showSelectChatGPTModels = false"
            />
          </van-popup>
          <van-field :label-width="200"
                     v-if="advancedConfig.ChatGptEnabled"
                     v-model="advancedConfig.ChatGptMaxTokens"
                     name="ChatGptMaxTokens"
                     :label="t('最大Token')"
                     type="digit"
                     :placeholder="t('请输入最大Token(1-2048)')"
                     :rules="[{ required: true, message: t('请输入最大Token(1-2048)'), validator: inputValidatorMaxToken }]"
          />
          <van-field :label-width="200"
                     v-if="advancedConfig.ChatGptEnabled"
                     v-model="advancedConfig.ChatGptTemperature"
                     name="ChatGptTemperature"
                     :label="t('随机性')"
                     type="number"
                     :placeholder="t('请输入随机性(0-1)')"
                     :rules="[{ required: true, message: t('请输入随机性(0-1)'), validator: inputValidatorTemperature }]"
          />
          <van-field :label-width="200"
                     v-if="advancedConfig.ChatGptEnabled"
                     v-model="advancedConfig.ChatGptTopP"
                     name="ChatGptTopP"
                     :label="t('核采样')"
                     type="number"
                     :placeholder="t('请输入核采样(0-1)')"
                     :rules="[{ required: true, message: t('请输入核采样(0-1)'), validator: inputValidatorTemperature }]"
          />
          <van-field :label-width="200"
                     v-if="advancedConfig.ChatGptEnabled"
                     v-model="advancedConfig.ChatGptPresencePenalty"
                     name="ChatGptPresencePenalty"
                     :label="t('重复惩罚')"
                     type="number"
                     :placeholder="t('请输入重复惩罚(0-1)')"
                     :rules="[{ required: true, message: t('请输入重复惩罚(0-1)'), validator: inputValidatorTemperature }]"
          />
          <van-field :label-width="200"
                     v-if="advancedConfig.ChatGptEnabled"
                     v-model="advancedConfig.ChatGptFrequencyPenalty"
                     name="ChatGptFrequencyPenalty"
                     :label="t('频繁惩罚')"
                     type="number"
                     :placeholder="t('请输入频繁惩罚(0-1)')"
                     :rules="[{ required: true, message: t('请输入频繁惩罚(0-1)'), validator: inputValidatorTemperature }]"
          />
          <van-field :label-width="200"
                     v-if="advancedConfig.ChatGptEnabled"
                     v-model="advancedConfig.ChatGptHistoryCount"
                     name="ChatGptHistoryCount"
                     :label="t('上文数量')"
                     type="digit"
                     :placeholder="t('请输入上文数量(1-20)')"
                     :rules="[{ required: true, message: t('请输入上文数量(1-20)'), validator: inputValidatorHistoryCount }]"
          />
          <van-field :label-width="200"
                     v-if="advancedConfig.ChatGptEnabled"
                     v-model="advancedConfig.ChatGptTimeout"
                     name="ChatGptTimeout"
                     :label="t('超时时间')"
                     type="digit"
                     :placeholder="t('请输入超时时间(1-600s)')"
                     :rules="[{ required: true, message: t('请输入超时时间(1-600s)'), validator: inputValidatorTimeout }]"
          />
          <van-field :label-width="200"
                     v-if="advancedConfig.ChatGptEnabled"
                     v-model="advancedConfig.ChatGptSystemPrompt"
                     name="ChatGptSystemPrompt"
                     :label="t('系统提示')"
                     :placeholder="t('请输入系统提示, 默认为You are a friendly AI created by OpenAI. You are very helpful and will answer any questions you ask.')"
                     type="textarea"
                     autosize
          />
          <van-field :label-width="200"
                     v-if="advancedConfig.ChatGptEnabled"
                     v-model="advancedConfig.ChatGptBaseUrl"
                     name="ChatGptBaseUrl"
                     :label="t('接口地址')"
                     :placeholder="t('请输入接口地址, 默认为官方地址')"
          />
          <van-field :label-width="200"
                     v-if="advancedConfig.ChatGptEnabled"
                     v-model="advancedConfig.ChatGptHttpProxy"
                     name="ChatGptHttpProxy"
                     :label="t('HTTP代理')"
                     :placeholder="t('请输入HTTP代理地址, 默认不代理')"
          />
        </van-cell-group>

        <h2 class="van-doc-demo-block__title" v-for="i in 3" :key="i">&nbsp;</h2>
      </div>

      <van-action-bar>
        <van-action-bar-icon icon="chat-o" :text="t('客服')"
                             @click="onClickService"
        />
        <!--    <van-action-bar-icon icon="shop-o" text="店铺"/>-->
        <van-action-bar-button color="#282828" type="warning" :text="t('重置')"/>
        <van-action-bar-button color="#3478f6" type="primary"
                               @click="onSubmitForm"
                               native-type="submit" :text="t('保存')"/>
      </van-action-bar>
    </van-form>

  </div>
</template>

<style lang="scss" scoped>
.editKeyboardGroupNameDialogDiv {
  width: 100%;
  height: 100px;
  //内容左右居中 上下居中
  display: flex;
  justify-content: center;
  align-items: center;
  //背景色#090909
  background-color: #282828;
  //内容上下排列
  flex-direction: column;
}

.editKeyboardGroupNameInput {
  width: 100%;
  height: 30px;
  border: none;
  background-color: #323232;
  border-radius: 5px;
  padding: 0 10px;
  font-size: 10px;
  color: #dfdfdf;
  //placeholder颜色
  &::placeholder {
    color: #707070;
  }
}

.van-col-keyboard {
  //内容左右居中 上下居中
  display: flex;
  justify-content: center;
  align-items: center;
  //圆角
  border-radius: 10px;
  //文字大小 10px
  font-size: 10px;
}

.page {
  display: flex;
  justify-content: center;
  //align-items: center;
  height: 100%;
  width: 100%;
  overflow-y: auto;

  .form {
    width: 100%;
    display: flex;
    justify-content: center;
    //align-items: center;
    height: 100%;

    .container {
      width: 100%;
    }
  }
}


.superLink {
  color: #76d571;
  cursor: pointer;
  font-size: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.superLinkDanger {
  color: #eb4d3d;
  cursor: pointer;
  font-size: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.superLinkInfo {
  color: #9e9e9e;
  cursor: pointer;
  font-size: 10px;
  display: flex;
  justify-content: center;
  align-items: center;

}

//inputAndClose中包含一个input和一个closeIcon，要求icon在input右上角
.inputAndClose {
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;

  .van-icon {
    position: absolute;
    right: 0;
    top: 0;
    color: #fff;
  }
}

//keyInput输入框
.keyInput {
  width: 100%;
  height: 100%;
  border: none;
  background-color: #9e9e9e;
  border-radius: 5px;
  padding: 0 10px;
  font-size: 10px;
  color: #dfdfdf;
  //placeholder颜色
  &::placeholder {
    color: #707070;
  }
}

//margin-left-14px
.margin-left-14px {
  margin-left: 14px;
}

//cursor-pointer
.cursor-pointer {
  cursor: pointer;
}
</style>
