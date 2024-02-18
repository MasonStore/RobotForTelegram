export namespace views {
	
	export class AddAccountReq {
	    token: string;
	
	    static createFrom(source: any = {}) {
	        return new AddAccountReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.token = source["token"];
	    }
	}
	export class AddAccountResp {
	    ok: boolean;
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new AddAccountResp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ok = source["ok"];
	        this.error = source["error"];
	    }
	}
	export class GetAccountsRespAccount {
	    id: number;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    token: string;
	    first_name: string;
	    last_name: string;
	    username: string;
	    running: boolean;
	
	    static createFrom(source: any = {}) {
	        return new GetAccountsRespAccount(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.token = source["token"];
	        this.first_name = source["first_name"];
	        this.last_name = source["last_name"];
	        this.username = source["username"];
	        this.running = source["running"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GetAccountsResp {
	    ok: boolean;
	    error: string;
	    accounts: GetAccountsRespAccount[];
	
	    static createFrom(source: any = {}) {
	        return new GetAccountsResp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ok = source["ok"];
	        this.error = source["error"];
	        this.accounts = this.convertValues(source["accounts"], GetAccountsRespAccount);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GetHomeRespAdvancedConfig {
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
	
	    static createFrom(source: any = {}) {
	        return new GetHomeRespAdvancedConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ChatGptEnabled = source["ChatGptEnabled"];
	        this.ChatGptApiKey = source["ChatGptApiKey"];
	        this.ChatGptHttpProxy = source["ChatGptHttpProxy"];
	        this.ChatGptModel = source["ChatGptModel"];
	        this.ChatGptSystemPrompt = source["ChatGptSystemPrompt"];
	        this.ChatGptBaseUrl = source["ChatGptBaseUrl"];
	        this.ChatGptTemperature = source["ChatGptTemperature"];
	        this.ChatGptMaxTokens = source["ChatGptMaxTokens"];
	        this.ChatGptTopP = source["ChatGptTopP"];
	        this.ChatGptPresencePenalty = source["ChatGptPresencePenalty"];
	        this.ChatGptFrequencyPenalty = source["ChatGptFrequencyPenalty"];
	        this.ChatGptHistoryCount = source["ChatGptHistoryCount"];
	        this.ChatGptTimeout = source["ChatGptTimeout"];
	    }
	}
	export class GetHomeRespCommand {
	    id: string;
	    command: string;
	    replyId: string;
	    replyName: string;
	    replyIsTemplate: string;
	    hookUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new GetHomeRespCommand(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.command = source["command"];
	        this.replyId = source["replyId"];
	        this.replyName = source["replyName"];
	        this.replyIsTemplate = source["replyIsTemplate"];
	        this.hookUrl = source["hookUrl"];
	    }
	}
	export class GetHomeRespButton {
	    Title: string;
	    Url: string;
	
	    static createFrom(source: any = {}) {
	        return new GetHomeRespButton(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Title = source["Title"];
	        this.Url = source["Url"];
	    }
	}
	export class GetHomeRespText {
	    Content: string;
	
	    static createFrom(source: any = {}) {
	        return new GetHomeRespText(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Content = source["Content"];
	    }
	}
	export class GetHomeRespPic {
	    Base64: string;
	
	    static createFrom(source: any = {}) {
	        return new GetHomeRespPic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Base64 = source["Base64"];
	    }
	}
	export class GetHomeRespGame {
	    GameShortName: string;
	    GameTitle: string;
	    Url: string;
	    ButtonText: string;
	
	    static createFrom(source: any = {}) {
	        return new GetHomeRespGame(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.GameShortName = source["GameShortName"];
	        this.GameTitle = source["GameTitle"];
	        this.Url = source["Url"];
	        this.ButtonText = source["ButtonText"];
	    }
	}
	export class GetHomeRespReply {
	    id: string;
	    name: string;
	    isTemplate: string;
	    messageType: string;
	    gameContent: GetHomeRespGame;
	    picContent: GetHomeRespPic;
	    textContent: GetHomeRespText;
	    expandType: string;
	    keyboardGroupId: string;
	    keyboardGroupName: string;
	    buttons: GetHomeRespButton[];
	
	    static createFrom(source: any = {}) {
	        return new GetHomeRespReply(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.isTemplate = source["isTemplate"];
	        this.messageType = source["messageType"];
	        this.gameContent = this.convertValues(source["gameContent"], GetHomeRespGame);
	        this.picContent = this.convertValues(source["picContent"], GetHomeRespPic);
	        this.textContent = this.convertValues(source["textContent"], GetHomeRespText);
	        this.expandType = source["expandType"];
	        this.keyboardGroupId = source["keyboardGroupId"];
	        this.keyboardGroupName = source["keyboardGroupName"];
	        this.buttons = this.convertValues(source["buttons"], GetHomeRespButton);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GetHomeRespKeyboardGroup {
	    id: string;
	    name: string;
	    keys: string[][];
	
	    static createFrom(source: any = {}) {
	        return new GetHomeRespKeyboardGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.keys = source["keys"];
	    }
	}
	export class GetHomeRespRobotInfo {
	    firstName: string;
	    lastName: string;
	    username: string;
	    token: string;
	
	    static createFrom(source: any = {}) {
	        return new GetHomeRespRobotInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.firstName = source["firstName"];
	        this.lastName = source["lastName"];
	        this.username = source["username"];
	        this.token = source["token"];
	    }
	}
	export class GetHomeRespData {
	    running: boolean;
	    robotInfo: GetHomeRespRobotInfo;
	    keyboardGroups: GetHomeRespKeyboardGroup[];
	    replies: GetHomeRespReply[];
	    commands: GetHomeRespCommand[];
	    advancedConfig: GetHomeRespAdvancedConfig;
	
	    static createFrom(source: any = {}) {
	        return new GetHomeRespData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.running = source["running"];
	        this.robotInfo = this.convertValues(source["robotInfo"], GetHomeRespRobotInfo);
	        this.keyboardGroups = this.convertValues(source["keyboardGroups"], GetHomeRespKeyboardGroup);
	        this.replies = this.convertValues(source["replies"], GetHomeRespReply);
	        this.commands = this.convertValues(source["commands"], GetHomeRespCommand);
	        this.advancedConfig = this.convertValues(source["advancedConfig"], GetHomeRespAdvancedConfig);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GetHomeResp {
	    ok: boolean;
	    error: string;
	    data: GetHomeRespData;
	
	    static createFrom(source: any = {}) {
	        return new GetHomeResp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ok = source["ok"];
	        this.error = source["error"];
	        this.data = this.convertValues(source["data"], GetHomeRespData);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	
	
	
	
	
	
	
	export class GetSelectedAccountResp {
	    ok: boolean;
	    error: string;
	    id: number;
	
	    static createFrom(source: any = {}) {
	        return new GetSelectedAccountResp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ok = source["ok"];
	        this.error = source["error"];
	        this.id = source["id"];
	    }
	}
	export class SelectOneAccountResp {
	    ok: boolean;
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new SelectOneAccountResp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ok = source["ok"];
	        this.error = source["error"];
	    }
	}

}

