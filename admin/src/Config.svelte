<script>
    import Scheduler from "./Scheduler.svelte";
</script>

<h2>Configuration</h2>
<br/>
<ul class="nav nav-tabs" id="myTab" role="tablist">
    <li class="nav-item">
        <a class="nav-link active" id="data-ex-tab1-tab" data-toggle="tab" href="#tab1" role="tab" aria-controls="tab1" aria-selected="true">Scheduler</a>
    </li>
    <li class="nav-item">
        <a class="nav-link" id="data-ex-tab2-tab" data-toggle="tab" href="#tab2" role="tab" aria-controls="tab2" aria-selected="false">IO-API</a>
    </li>
</ul>

        <!-- Tab panes -->
<div class="tab-content">
    <div class="tab-pane active" id="tab1" role="tabpanel" aria-labelledby="data-ex-tab1-tab">
        <Scheduler />
    </div>
    <div class="tab-pane" id="tab2" role="tabpanel" aria-labelledby="data-ex-tab2-tab">more to come</div>
</div>
