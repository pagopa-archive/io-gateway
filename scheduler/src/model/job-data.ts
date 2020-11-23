/**
 * Represents a simple JobData that could be scheduled using
 * a cron like time expression.
 */
class JobData {

    public jobName: string;
    public time: string;
    public action: string;

    constructor(jobName: string, time: string, action: string) {
        this.jobName = jobName;
        this.time = time;
        this.action = action;
    }

}

export default JobData;
