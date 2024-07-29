import { createI18n } from "vue-i18n";

const messages = {
    en: {
      message: {
        mainpage: 'Home',
        problemset: 'ProblemSet',
        status: 'Status',
        discussion: 'Discussion',
        about: 'About',
        usermanagement: 'User Management',
        problemmanagement: 'Problem Management',
        lang: 'Language',
        login: 'Login',
        logout: "Logout",
        register: 'Register',
        forget: 'Forget Password?',
        username: 'Username',
        password: 'Password',
        confirmPwd: 'Confirm Password',
        vCode: 'Verification Code',
        profile: 'Profile',
        resetpwd: 'Reset Password',
        createDiscussion: 'Create Discussion',
        nopermission: 'No Permission',
        loading: 'Loading...',
        notice: 'Notice',
        searchProblem: 'Problem Search',
        searchUser: 'User Search',
        problem: 'Problem',
        difficulty: 'Difficulty',
        problemId: 'Problem ID',
        when: 'When',
        userId: 'User ID',
        result: 'Result',
        nodata: 'No Data',
        nologin: 'Please login and check again!',
        title: 'Title',
        submit: 'Submit',
        discussionOwner: 'Owner',
        comments: 'Comments',
        delete: 'Delete',
        banUser: 'Block this user',
        unbanUser: 'Unblock this user',
        demoteUser: 'Demote this user',
        promoteUser: 'Promote this user',
        email: 'Email',
        role: 'Role',
        timeOfRegister: 'Join Time',
        operation: 'Operation',
        admin: 'Admin',
        regularUser: 'Regular User',
        isBan: 'Blocked',
        normal: 'Normal',
        timeLimit: 'Time Limit',
        memoryLimit: 'Memory Limit',
        displayInputCase: 'Input',
        displayOutputCase: 'Output',
        testcases: "Testcases",
        input: "Input",
        output: "Output",
        addTestCase: "Add Testcase",
        deleteTestCase: "Delete Testcase",
        cancel: "Cancel",
        save: "Save",
        edit: "Edit",
        notify: "Notify",
        ok: "OK",
        success: "Success",
        failed: "Failed, Please try again",
        tokenCheckfailed: "Token Check Failed, Please login again",
        formCheckfailed: "Make sure the form content is not empty",
        formRuleCheckfailed: "Fill out the form according to the corresponding rules",
        checkEmail: "Please check your email",
        usernameRequired: "Username is required.",
        emailRequired: "Email is required.",
        vCodeRequired: "Verification Code is required.",
        passwordRequired: "Password is required.",
        confirmPwdRequired: "Confirm Password is required.",
        validEmail: "Please enter a valid email address.",
        vCodeLength: "Verification Code must be 4 digits.",
        passwordLength: "Password must be 8-20 characters.",
        confirmPwdLength: "Confirm Password must be 8-20 characters.",
        back: "Back",
      },
    },
    zh_CN: {
      message: {
        mainpage: '主页',
        problemset: '题目',
        status: '状态',
        discussion: '讨论',
        about: '关于',
        usermanagement: '用户管理',
        problemmanagement: '题目管理',
        lang: '语言',
        login: '登录',
        logout: "登出",
        register: '注册',
        forget: '忘记密码？',
        username: '用户名',
        password: '密码',
        confirmPwd: '确认密码',
        vCode: '验证码',
        profile: '个人资料',
        resetpwd: '重置密码',
        createDiscussion: '创建讨论',
        nopermission: '没有权限',
        loading: '加载中...',
        notice: '通知',
        searchProblem: '题目搜索',
        searchUser: '用户搜索',
        problem: '题目',
        difficulty: '难度',
        problemId: '题目ID',
        when: '时间',
        userId: '用户ID',
        result: '结果',
        nodata: '没有数据',
        nologin: '请登录后再试！',
        title: '标题',
        submit: '提交',
        discussionOwner: '创建者',
        comments: '评论',
        delete: '删除',
        banUser: '封禁此用户',
        unbanUser: '解封此用户',
        demoteUser: '降级此用户',
        promoteUser: '升级此用户',
        email: '邮箱',
        role: '角色',
        timeOfRegister: '注册时间',
        operation: '操作',
        admin: '管理员',
        regularUser: '普通用户',
        isBan: '封禁',
        normal: '正常',
        timeLimit: '时间限制',
        memoryLimit: '内存限制',
        displayInputCase: '输入样例',
        displayOutputCase: '输出样例',
        testcases: "测试样例",
        input: "输入",
        output: "输出",
        addTestCase: "添加测试样例",
        deleteTestCase: "删除测试样例",
        cancel: "取消",
        save: "保存",
        edit: "编辑",
        notify: "通知",
        ok: "确定",
        success: "成功",
        failed: "失败，请重试",
        tokenCheckfailed: "Token验证失败，请重新登录",
        formCheckfailed: "请确保表单内容不为空",
        formRuleCheckfailed: "请按照对应规则填写表单",
        checkEmail: "请检查您的邮箱",
        usernameRequired: "用户名是必填项。",
        emailRequired: "邮箱是必填项。",
        vCodeRequired: "验证码是必填项。",
        passwordRequired: "密码是必填项。",
        confirmPwdRequired: "确认密码是必填项。",
        validEmail: "请输入有效的邮箱地址。",
        vCodeLength: "验证码必须是4位数字。",
        passwordLength: "密码必须是8-20位字符。",
        confirmPwdLength: "确认密码必须是8-20位字符。",
        back: "返回",
      },
    },
    ja: {
      message: {
        mainpage: 'ホームページ',
        problemset: '問題セット',
        status: 'ステータス',
        discussion: 'ディスカッション',
        about: 'アバウト',
        usermanagement: 'ユーザー管理',
        problemmanagement: '問題管理',
        lang: '言語',
        login: 'ログイン',
        logout: "ログアウト",
        register: '登録ユーザー',
        forget: '忘れたのか？',
        username: 'ユーザー名',
        password: 'パスワード',
        vCode: '確認コード',
        confirmPwd: 'パスワード確認',
        profile: 'プロフィール',
        resetpwd: 'パスワードをリセット',
        createDiscussion: 'ディスカッションを作成',
        nopermission: '権限がありません',
        loading: 'ロード中...',
        notice: 'お知らせ',
        searchProblem: '問題検索',
        searchUser: 'ユーザー検索',
        problem: '問題',
        difficulty: '難易度',
        problemId: '問題ID',
        when: 'いつ',
        userId: 'ユーザーID',
        result: '結果',
        nodata: 'データがありません',
        nologin: 'ログインしてください！',
        title: 'タイトル',
        submit: '提出',
        discussionOwner: '作成者',
        comments: 'コメント',
        delete: '削除',
        banUser: 'このユーザーをブロック',
        unbanUser: 'このユーザーをブロック解除',
        demoteUser: 'このユーザーを降格',
        promoteUser: 'このユーザーを昇格',
        email: 'メール',
        role: '役割',
        timeOfRegister: '登録日時',
        operation: '操作',
        admin: '管理者',
        regularUser: '一般ユーザー',
        isBan: 'ブロック',
        normal: '正常',
        timeLimit: '時間制限',
        memoryLimit: 'メモリ制限',
        displayInputCase: '入力例を',
        displayOutputCase: '出力例を',
        testcases: "テストケース",
        input: "入力",
        output: "出力",
        addTestCase: "テストケースを追加",
        deleteTestCase: "テストケースを削除",
        cancel: "キャンセル",
        save: "保存",
        edit: "編集",
        notify: "通知",
        ok: "OK",
        success: "成功",
        failed: "失敗，再試してください",
        tokenCheckfailed: "トークン認証に失敗しました，再ログインしてください",
        formCheckfailed: "すべてのフォームを入力してください",
        formRuleCheckfailed: "対応するルールに従ってフォームに記入する",
        checkEmail: "メールを確認してください",
        usernameRequired: "ユーザー名は必須です。",
        emailRequired: "メールは必須です。",
        vCodeRequired: "確認コードは必須です。",
        passwordRequired: "パスワードは必須です。",
        confirmPwdRequired: "確認パスワードは必須です。",
        validEmail: "有効なメールアドレスを入力してください。",
        vCodeLength: "確認コードは4桁の数字です。",
        passwordLength: "パスワードは8-20文字です。",
        confirmPwdLength: "確認パスワードは8-20文字です。",
        back: "戻る",
      },
    },
    zh_TW: {
      message: {
        mainpage: '主頁',
        problemset: '題目',
        status: '狀態',
        discussion: '討論',
        about: '關於',
        usermanagement: '用戶管理',
        problemmanagement: '題目管理',
        lang: '語言',
        login: '登錄',
        logout: "登出",
        register: '註冊',
        forget: '忘記密碼？',
        username: '用戶名',
        password: '密碼',
        confirmPwd: '確認密碼',
        vCode: '驗證碼',
        profile: '個人資料',
        resetpwd: '重設密碼',
        createDiscussion: '創建討論',
        nopermission: '沒有權限',
        loading: '加載中...',
        notice: '通知',
        searchProblem: '題目搜索',
        searchUser: '用戶搜索',
        problem: '題目',
        difficulty: '難度',
        problemId: '題目ID',
        when: '時間',
        userId: '用戶ID',
        result: '結果',
        nodata: '沒有數據',
        nologin: '請先登錄！',
        title: '標題',
        submit: '提交',
        discussionOwner: '創建者',
        comments: '評論',
        delete: '刪除',
        banUser: '封禁此用戶',
        unbanUser: '解封此用戶',
        demoteUser: '降級此用戶',
        promoteUser: '升級此用戶',
        email: '郵箱',
        role: '角色',
        timeOfRegister: '註冊時間',
        operation: '操作',
        admin: '管理員',
        regularUser: '普通用戶',
        isBan: '封禁',
        normal: '正常',
        timeLimit: '時間限制',
        memoryLimit: '內存限制',
        displayInputCase: '輸入用例',
        displayOutputCase: '輸出用例',
        testcases: "測試用例",
        input: "輸入",
        output: "輸出",
        addTestCase: "添加測試用例",
        deleteTestCase: "刪除測試用例",
        cancel: "取消",
        save: "保存",
        edit: "編輯",
        notify: "通知",
        ok: "確定",
        success: "成功",
        failed: "失敗，請再試一次",
        tokenCheckfailed: "token認證失敗，請重新登錄",
        formCheckfailed: "請填寫所有表單",
        formRuleCheckfailed: "請按照對應規則填寫表單",
        checkEmail: "請確認郵箱",
        usernameRequired: "用戶名是必填項。",
        emailRequired: "郵箱是必填項。",
        vCodeRequired: "確認碼是必填項。",
        passwordRequired: "密碼是必填項。",
        confirmPwdRequired: "確認密碼是必填項。",
        validEmail: "請輸入有效的郵箱地址。",
        vCodeLength: "確認碼是4位數字。",
        passwordLength: "密碼長度為8-20個字符。",
        confirmPwdLength: "確認密碼長度為8-20個字符。",
        back: "返回",
      },
    },
};

export const i18n = createI18n({
    legacy: false,
    locale: 'en', 
    fallbackLocale: 'en',
    messages,
})