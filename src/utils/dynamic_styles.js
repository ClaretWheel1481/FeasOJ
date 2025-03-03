// 根据结果不同显示不同颜色
export const getResultStyle = (result) => {
    switch (result) {
        case 'Compile Failed':
            return 'color: red; font-weight: bold;';
        case 'Time Limit Exceeded':
            return 'color: red; font-weight: bold;';
        case 'Memory Limit Exceeded':
            return 'color: red; font-weight: bold;';
        case 'Internal Error':
            return 'color: red; font-weight: bold';
        case 'Success':
            return 'color: green; font-weight: bold;';
        case 'Failed':
            return 'color: orange; font-weight: bold;';
        case 'Wrong Answer':
            return 'color: orange; font-weight: bold;';
        default:
            return '';
    }
};

// 根据题目难度显示不同字体
export const difficultyColor = (difficulty) => {
    switch (difficulty) {
        case '简单':
            return 'font-weight: bold;color: green;';
        case '中等':
            return 'font-weight: bold;color: orange;';
        case '困难':
            return 'font-weight: bold;color: red;';
        default:
            return '';
    }
};
export const difficultyLang = (difficulty) => {
    switch (difficulty) {
        case '简单':
            return 'message.easy';
        case '中等':
            return 'message.medium';
        case '困难':
            return 'message.hard';
        default:
            return '';
    }
}