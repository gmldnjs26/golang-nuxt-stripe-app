const log4js = require('log4js')

const d = new Date()

/** 日付を文字列にフォーマットする */
const directoryFormatted = `${d.getFullYear()}${(d.getMonth() + 1)
  .toString()
  .padStart(2, '0')}${d.getDate().toString().padStart(2, '0')}`.replace(
  /\n|\r/g,
  ''
)

const patternFormatted = `-${d.getFullYear()}-${(d.getMonth() + 1)
  .toString()
  .padStart(2, '0')}-${d.getDate().toString().padStart(2, '0')}`.replace(
  /\n|\r/g,
  ''
)

log4js.configure({
  appenders: {
    console: {
      type: 'console',
    },
    system: {
      type: 'dateFile',
      filename: `logs/${directoryFormatted}/heewon_NodeServer.log${patternFormatted}`,
    },
  },
  categories: {
    default: {
      appenders: ['console', 'system'],
      level: 'ALL',
    },
  },
})

const systemLogger = log4js.getLogger('system')

module.exports = { systemLogger }
