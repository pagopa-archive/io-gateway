import chai from 'chai';
import SchedulerConfigurator from "../src/config/scheduler-configurator";
import JobData from "../src/model/job-data";
import fs from 'fs';

const aJob = new JobData('testJob',"*/1 * * * *",'echo this is a message');
const configurator  = new SchedulerConfigurator();

// set the env variable to force the configurator to use a custom config filename.
process.env.IO_SDK_SCHEDULER_CONFIG = process.env.HOME+'/io-sdk-scheduler-config-test.json';

describe("SchedulerConfigurator",function(){
    describe("persistJobs", function(){
        it('should persist jobs to the configured config file', async function () {

            const jobArray = [];
            jobArray.push(aJob);
            configurator.persistJobs(jobArray);
            chai.expect(true).to.equal(fs.existsSync(configurator.getConfigFilename()));
        });
    })
});

describe("SchedulerConfigurator",function(){
    describe("retrievePersistedJobs", function(){
        it('should retrieve 1 persisted job from config file', async function () {
            const jobArray = configurator.retrievePersistedJobs();

            console.log(JSON.stringify(jobArray));
            chai.expect(1).to.equal(jobArray.length);

            const persistedJob = jobArray[0];
            chai.expect(aJob.jobName).to.equal(persistedJob.jobName);
        });
    })
});
