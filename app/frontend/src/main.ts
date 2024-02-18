import {createApp} from "vue";
import {createPinia} from "pinia";
import App from "./App.vue";
import router from "./router";
import i18n from "./i18n";
import {
    Button,
    ActionBar,
    ActionBarButton,
    ActionBarIcon,
    Cell,
    CellGroup,
    Collapse,
    CollapseItem,
    ConfigProvider,
    Icon,
    Notify,
    setNotifyDefaultOptions,
    Switch,
    Toast,
    Col,
    Row,
    Divider,
    Dialog,
    Field,
    Form,
    SubmitBar,
    Radio,
    RadioGroup,
    Uploader,
    Image as VanImage,
    ImagePreview,
    Picker,
    Popup,
    Stepper,
    Checkbox,
    NavBar,
} from "vant";
import "vant/lib/index.css";

import "./style.scss";

// import "./assets/main.css";

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(i18n);
app.use(Toast);
app.use(Notify);
app.use(ConfigProvider);
app.use(Icon);
app.use(ActionBar);
app.use(ActionBarIcon);
app.use(ActionBarButton);
app.use(Collapse);
app.use(CollapseItem);
app.use(Cell);
app.use(CellGroup);
app.use(Switch);
app.use(Col);
app.use(Row);
app.use(Divider);
app.use(Dialog);
app.use(Field);
app.use(Form);
app.use(SubmitBar);
app.use(Radio);
app.use(RadioGroup);
app.use(Uploader);
app.use(Button);
app.use(VanImage);
app.use(ImagePreview);
app.use(Picker);
app.use(Popup);
app.use(Stepper);
app.use(Checkbox);
app.use(NavBar);
setNotifyDefaultOptions({
    className: "rft-notify",
})
app.mount("#app");
