<script>
    import { onMount } from "svelte";

    let base = "http://localhost:3100/scheduler";

    let state = {};
    let scheduledJobs = [];

    async function getScheduledJobs() {
        console.log("fetching running jobs from scheduler");
        fetch(base, {
            method: "GET",
            headers: { "Content-Type": "application/json" }
        })
            .then(async res => {
                if (res.ok) {
                    scheduledJobs = await res.json();
                }
            })
            .catch(err => {
                state = { error: err };
            });
    }

    async function deleteJob(jobName) {
        return fetch(base, {
            method: "DELETE",
            body: JSON.stringify({jobName: jobName}),
            headers: { "Content-Type": "application/json" }
        })
            .then(async res => {
                if (res.ok) {
                    state = { result: 'Job '+jobName+' removed' };
                    await getScheduledJobs();
                }
            })
            .catch(err => {
                console.log(JSON.stringify(err));
                state = { error: err };
            });
    }

    function confirmRemoval(jobName) {
        console.log('called with ', jobName);
        if(!confirm("Are you sure you want to remove job "+jobName+" ?")) {
            return;
        }
        deleteJob(jobName);
    }

    onMount(() => getScheduledJobs());
</script>
<br />
<div>
    <table class="table">
        <thead>
        <tr>
            <th scope="col">Job</th>
            <th scope="col">Time Expression</th>
            <th scope="col">Command</th>
            <th scope="col">Status</th>
            <th scope="col">Action</th>
        </tr>
        </thead>
        <tbody>
            {#each scheduledJobs as rec}
            <tr>
                <td>{rec.job.jobName}</td>
                <td>{rec.job.time}</td>
                <td>{rec.job.action}</td>
                <td>{rec.status}</td>
                <td>
                    <button on:click={() => confirmRemoval(rec.job.jobName)} >
                        <i class="fa fa-trash" />
                    </button>
                </td>
            </tr>
            {:else}
            <tr>
                <td colspan="5">Info: no scheduled jobs available</td>
            </tr>
            {/each}
        </tbody>
    </table>
</div>

{#if state.result}
<div class="alert alert-success" role="alert">
    <pre>{state.result}</pre>
</div>
{:else if state.error}
<div class="alert alert-danger" role="alert">
    {state.error}
</div>
{/if}
