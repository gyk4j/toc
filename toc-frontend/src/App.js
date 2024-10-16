import React from 'react';
import PreTocCheckerDialog from './components/PreTocCheckerDialog';
import TocEndedDialog from './components/TocEndedDialog';
import StatusBar from './components/StatusBar';
import ToolBar from './components/ToolBar';
import TabbedPane from './components/TabbedPane';
import * as ApiService from './services/ApiService'

export default class App extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      name: "TOC",
      showPreTocCheckerDialog: true,
      preTocCheckerFileStatus: 'Completed',
      preTocCheckerWebStatus: 'Failed',
      preTocCheckerMailStatus: 'Running',
      preTocCheckerDatabaseStatus: 'Pending',
      preTocCheckerActionText: "Start full backup",
      showTocEndedDialog: false,
      statusbarIsMain: Math.random() > 0.5,
      statusbarIsOnline: Math.random() > 0.5,
      statusbarInterval: '1 minute',
      statusbarCountdown: -1, // in seconds
      statusbarStartTime: null,
      statusbarEndTime: null,
      toolbar: null,
      tabsStatisticsTotal: 100, // kb
      tabsStatisticsTransferred: 0, // kb
      tabsStatisticsDuration: Math.floor(Math.random() * (60 * 60)),
      tabsLogsDatabase: [
        "INFO: Line 1",
        "DEBUG: Line 2",
        "TRACE: Line 3",
        "ERROR: Line 4",
        "WARN: Line 5",
        "FATAL: Line 6"
      ],
      tabsLogsEmail: [
        "INFO: Line 1",
        "DEBUG: Line 2",
        "TRACE: Line 3",
        "ERROR: Line 4",
        "WARN: Line 5",
        "FATAL: Line 6"
      ],
      tabsLogsFile: [
        "INFO: Line 1",
        "DEBUG: Line 2",
        "TRACE: Line 3",
        "ERROR: Line 4",
        "WARN: Line 5",
        "FATAL: Line 6"
      ],
      tabsLogsWeb: [
        "INFO: Line 1",
        "DEBUG: Line 2",
        "TRACE: Line 3",
        "ERROR: Line 4",
        "WARN: Line 5",
        "FATAL: Line 6"
      ],
      tabsProgressBackups: [
        {
          time: new Date(),
          database: "Backup created",
          email: "Restoration completed",
          file: "Backup failed",
          web: "Restoration failed"
        },
        {
          time: new Date(),
          database: "Backup created",
          email: "Transferring to main controller",
          file: "Transferring to file server",
          web: "Restoring backup"
        }
      ],
      tabsQuotaQuotas: [
        {
          directory: "C:\\Windows",
          total: 100,
          used: Math.floor(Math.random() * 101)
        },
        {
          directory: "C:\\Windows",
          total: 100,
          used: Math.floor(Math.random() * 101)
        },
        {
          directory: "C:\\Windows",
          total: 100,
          used: Math.floor(Math.random() * 101)
        },
        {
          directory: "C:\\Windows\\Temp",
          total: 100,
          used: Math.floor(Math.random() * 101)
        },
        {
          directory: "C:\\Users",
          total: 100,
          used: Math.floor(Math.random() * 101)
        }
      ]
    };

    this.closePreTocChecker = this.closePreTocChecker.bind(this);
    this.toggleShowTocEnded = this.toggleShowTocEnded.bind(this);
    this.changeIntervalOnChange = this.changeIntervalOnChange.bind(this);
    this.startOnClick = this.startOnClick.bind(this);
    this.sendCompletedOnClick = this.sendCompletedOnClick.bind(this);
    this.sendDeltaNowOnClick = this.sendDeltaNowOnClick.bind(this);
    this.sendLastDeltaOnClick = this.sendLastDeltaOnClick.bind(this);
    this.checkQuotaOnClick = this.checkQuotaOnClick.bind(this);
    this.exportLogsOnClick = this.exportLogsOnClick.bind(this);
    this.archiveOnClick = this.archiveOnClick.bind(this);

    this.onTimer = this.onTimer.bind(this);

    this.timer = null;
    this.timerCountDown = null;
  }

  render() {
    
    return (
      <div>
        <PreTocCheckerDialog
          showPreTocCheckerDialog={this.state.showPreTocCheckerDialog}
          closePreTocChecker={this.closePreTocChecker}
          fileStatus={this.state.preTocCheckerFileStatus}
          webStatus={this.state.preTocCheckerWebStatus}
          mailStatus={this.state.preTocCheckerMailStatus}
          databaseStatus={this.state.preTocCheckerDatabaseStatus}
          actionText={this.state.preTocCheckerActionText}
        /> 
        
        <TocEndedDialog
           showTocEndedDialog={this.state.showTocEndedDialog}
           toggleShowTocEnded={this.toggleShowTocEnded}
           statusbarEndTime={this.state.statusbarEndTime}
           actionText='Close'
        />

        <StatusBar
          isMain={this.state.statusbarIsMain} 
          isOnline={this.state.statusbarIsOnline}
          interval={this.state.statusbarInterval} 
          changeIntervalOnChange={this.changeIntervalOnChange}
          countdown={this.state.statusbarCountdown}
          startTime={this.state.statusbarStartTime}
          endTime={this.state.statusbarEndTime}
        />
        <ToolBar
          transferState={!(this.state.statusbarStartTime == null || this.state.statusbarStartTime === 'Not started')? 'STARTED': 'NOT_STARTED'}
          startOnClick={this.startOnClick}
          sendCompletedOnClick={this.sendCompletedOnClick}
          sendDeltaNowOnClick={this.sendDeltaNowOnClick}
          sendLastDeltaOnClick={this.sendLastDeltaOnClick}
          syncStatusOnClick={this.syncStatusOnClick}
          checkQuotaOnClick={this.checkQuotaOnClick}
          exportLogsOnClick={this.exportLogsOnClick}
          archiveOnClick={this.archiveOnClick} 
        />
        <TabbedPane 
          statisticsTotal={this.state.tabsStatisticsTotal}
          statisticsTransferred={this.state.tabsStatisticsTransferred}
          statisticsDuration={this.state.tabsStatisticsDuration}
          logsDatabase={this.state.tabsLogsDatabase}
          logsEmail={this.state.tabsLogsEmail}
          logsFile={this.state.tabsLogsFile}
          logsWeb={this.state.tabsLogsWeb}
          progressBackups={this.state.tabsProgressBackups}
          quotaQuotas={this.state.tabsQuotaQuotas} 
        />
      </div>
    );
  }

  // Event handlers
  closePreTocChecker = ( e ) => {
    //let show = this.state.showPreTocCheckerModal
    this.setState({ showPreTocCheckerDialog: false })
  }

  toggleShowTocEnded = ( e ) => {
    let show = this.state.showTocEndedDialog
    this.setState({ showTocEndedDialog: !show })
  }

  changeIntervalOnChange = ( event ) => {
    //console.log(event);
    this.setState({statusbarInterval: event.target.text});
  }

  startOnClick = () => {
    // this.timer = setInterval(this.onTimer, 100);
    // this.setState({statusbarStartTime: new Date() });

    let interval = this.state.statusbarInterval.toLowerCase()
    let intervalInSec = 0
    let amount = parseInt(interval)
    
    if(interval.endsWith("minutes") || interval.endsWith("minute"))
      intervalInSec = amount * 60
    else if(interval.endsWith("hours") || interval.endsWith("hour"))
      intervalInSec = amount * 3600
    else if(interval === "daily")
      intervalInSec = 86400
    else if(interval === "weekly")
      intervalInSec = 604800
    else
      intervalInSec = 0

    console.log("interval = '" + interval + "', intervalInSec = " + intervalInSec + " seconds")

    this.resetCountDown = () => {
      let countdown = this.state.statusbarIntervalInSec
      this.setState({ statusbarCountdown: countdown })
    }
    this.resetCountDown = this.resetCountDown.bind(this)

    this.onBackup = () => {
      let now = new Date()
      ApiService.newBackup()
      this.resetCountDown()
    }
    this.onBackup = this.onBackup.bind(this)

    // Refresh timer countdown display every second
    this.onTimerRefresh = () => {
      let countdown = this.state.statusbarCountdown

      if(isNaN(countdown) || countdown < 0){
        clearInterval(this.timerRefresh)
        this.timerRefresh = null
        console.log("Stopped timer due to non-numeric/negative countdown")
      }
      else {
        countdown-- // assumes it is triggered every second
        this.setState({ statusbarCountdown: countdown }, () => {
          if(this.state.statusbarCountdown === 0){
            this.onBackup()
          }
        })
      }
    }
    this.onTimerRefresh = this.onTimerRefresh.bind(this)

    this.onIntervalUpdated = () => {
      let intervalInSec = this.state.statusbarIntervalInSec
      if(Number.isInteger(intervalInSec)){
        console.log("onIntervalUpdated: Starting")
        this.setState({ statusbarStartTime: new Date() })
        let countdown = intervalInSec
        this.setState({ statusbarCountdown: countdown })

        if(this.timerRefresh != null)
          clearInterval(this.timerRefresh)

        this.timerRefresh = setInterval(this.onTimerRefresh, 1000)
      }
      else{
        alert("intervalInSec is NaN! " + intervalInSec)
        this.sendLastDeltaOnClick()
      }
    }
    this.onIntervalUpdated = this.onIntervalUpdated.bind(this)

    this.setState({ statusbarIntervalInSec: intervalInSec }, this.onIntervalUpdated);
  }

  sendCompletedOnClick = () => {
    let backup = {
      id : 0,
      snapshots : [],
      time : "1999-09-19T19:19:19.999+08:00"
    }
   
    ApiService.newRestoration(backup)
  }

  sendDeltaNowOnClick = () => {
    ApiService.newBackup()
  }

  sendLastDeltaOnClick = () => {

    if(this.timerRefresh != null){
      clearInterval(this.timerRefresh)
      this.timerRefresh = null
      console.log("Cleared refresh timer")
    }
    
    ApiService.newBackup()

    this.setState({statusbarEndTime: new Date() }, () => {
      this.toggleShowTocEnded()
    });
  }

  syncStatusOnClick = () => {
    ApiService.newSynchronization()
  }

  checkQuotaOnClick = () => {
    ApiService.getQuotas()
  }

  exportLogsOnClick = () => {
    ApiService.newLogs()
  }

  archiveOnClick = () => {
    ApiService.newArchive()
  }

  onTimer = () => {
    const stopTimer = () => {
      clearInterval(this.timer);
      this.timer = null;
    }

    stopTimer.bind(this);

    let current = this.state.tabsStatisticsTransferred;
    
    if(isNaN(current)){
      stopTimer();
    }
    else {
      var progress = ++current;

      if(isNaN(progress)){
        stopTimer();
      }
      else {
        this.setState({tabsStatisticsTransferred: progress }, ()=> {
          if(this.state.tabsStatisticsTransferred === this.state.tabsStatisticsTotal){
            stopTimer();
          }
        });
      }
    }
  }
}