import chai from 'chai';
import assert from 'assert';
import Storage from "../src/store/storage";
import JobData from "../src/model/job-data";
import CommandLineCronJob from "../src/job/command-line-cron-job";

const store = new Storage();
const aJob = new JobData('testJob',"*/1 * * * *",'echo this is a message');
const cronJob = new CommandLineCronJob(aJob);

describe("Storage",function(){
    describe("save", function(){
        it('should store a job into the internal Map', function () {
            store.save(cronJob, false);
            chai.expect(1).to.equal(store.getJobs().length);
        });
    })
});

describe("getJob", function(){
    it('should return a job object corresponding to the jobKey', function () {
        const foundJob: CommandLineCronJob = store.getJob(cronJob.getJobKey());
        chai.expect(cronJob.getJobKey()).to.equal(foundJob.getJobKey());
    });
})
