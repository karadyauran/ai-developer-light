    footprintDisplay: {
        fontSize: '18px',
        color: '#333',
        margin: '10px 0'
    },
    resultDisplay: {
        fontSize: '16px',
        color: '#555',
        margin: '10px 0'
    }
};

export function applyStyles() {
    const footprintElement = document.getElementById('footprint-display');
    const resultElement = document.getElementById('result-display');
    Object.assign(footprintElement.style, styles.footprintDisplay);
    Object.assign(resultElement.style, styles.resultDisplay);