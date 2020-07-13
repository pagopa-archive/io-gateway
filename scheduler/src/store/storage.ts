import CommandLineCronJob from "../job/command-line-cron-job";
import SchedulingResult from "../model/scheduling-result";
import SchedulerConfigurator from "../config/scheduler-configurator";
import JobData from "../model/job-data";

class Storage {

    private configurator: SchedulerConfigurator;
    private scheduledTasks = new Map();

    constructor() {
        this.configurator = new SchedulerConfigurator();
    }

    /**
     * Add a JobData into the JobStorage.
     *
     * @param aJob
     * @param persist if true force store the job also in the persisted configuration.
     */
    save(aJob: CommandLineCronJob, persist: boolean){
        this.scheduledTasks.set(aJob.getJobKey(), aJob);

        if(persist) {
            this.configurator.persistJobs(this.getJobsToPersist());
        }
    }

    /**
     * Return hte scheduled JobData Mapping the name.
     *
     * @param jobName
     */
    getJob(jobName): CommandLineCronJob {

        if(this.scheduledTasks.has(jobName)) {
            return this.scheduledTasks.get(jobName);
        }

        return null;
    }

    /**
     * Remove the scheduled JobData Mapping the name.
     *
     * @param jobName
     */
    removeJob(jobName): void {

        if(this.scheduledTasks.has(jobName)) {
            this.scheduledTasks.delete(jobName);
            this.configurator.persistJobs(this.getJobsToPersist());
        }
    }

    getJobs(): Array<SchedulingResult> {
        const result = [];
        this.scheduledTasks.forEach((job: CommandLineCronJob) => {
            result.push(job.getSchedulingResult());
        })
        return result;
    }

    private getJobsToPersist(): Array<JobData> {
        const result = [];
        this.scheduledTasks.forEach((job: CommandLineCronJob) => {
            result.push(job.getSchedulingResult().job);
        })
        return result;
    }

}
export default Storage;
