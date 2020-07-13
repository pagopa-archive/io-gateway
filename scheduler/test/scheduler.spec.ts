import chai from 'chai';
import assert from 'assert';

import JobData from "../src/model/job-data";
import CommandLineCronJob from "../src/job/command-line-cron-job";
import Scheduler from "../src/scheduler/scheduler";
import SchedulingResult from "../src/model/scheduling-result";

const scheduler = new Scheduler();
const aJob = new JobData('testJob',"*/1 * * * *",'echo this is a message');


describe("Scheduler",function(){
    describe("scheduleJob", function(){
        it('should schedule a Job and return a scheduled instance', function () {
            const schedulerResult: SchedulingResult = scheduler.scheduleJob(aJob, false);
            chai.expect("scheduled").to.equal(schedulerResult.status);
        });
    })
});

describe("Scheduler",function(){
    describe("getScheduledJobs", function(){
        it('should return an array od scheduled job of length 1', function () {
            const result: SchedulingResult[] = scheduler.getScheduledJobs();
            chai.expect(1).to.equal(result.length);
        });
    })
});

describe("Scheduler",function(){
    describe("removeScheduledJob", function(){
        it('should remove the requested Job and end the test', function () {
            scheduler.removeScheduledJob(aJob.jobName);
        });
    })
});
