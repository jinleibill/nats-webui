class Utils {
    formatData = (time) => {
        return (new Date(time)).toISOString().replace(/T/, " ").replace(/\.\w+/, '')
    }
}

export default new Utils;