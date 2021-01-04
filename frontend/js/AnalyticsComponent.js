
const NameComponentKey = "rb-name-component"

var NameComponent = {
    props: ['title', 'avatar_src'],
    template: `
    <h3 style="display: flex;justify-content: center;align-items: center;padding-left: 8px;font-family: 'Fira Code', monospace;color: #f5f3f1">
        {{ title }}
    </h3>      
    `
}

const ProjectHeaderComponentKey = "rb-project-header"

var ProjectHeaderComponent = {
    props: ['project'],
    template: `
    <div slot="header">
        <vs-row vs-justify="flex-start">
            <vs-avatar
                v-if="project.imageUrl != null"
                :src="project.imageUrl"
                size="large" 
            ></vs-avatar>
            <rb-name-component
                v-bind:title="project.name"
            ></rb-name-component>
        </vs-row>
    </div>
    `
}

const ChartComponentKey = "rb-chart-component"

var ChartComponent = {
    props: ['title','canvasId'],
    template: `
    <vs-card style="background-color: #313131" actionable class="cardx">
        <div slot="header">
            <rb-name-component
                v-bind:title="title"
            ></rb-name-component>
        </div>
        <div class="chart-container" style="position: relative; height:160px; width: inherit; background-color: #313131">
            <canvas :id="canvasId"></canvas>
        </div>
        <slot name="bottom-particle"></slot>
    </vs-card>
    `
}

const ProgressBarKey = "rb-progress-bar"

var ProgressBar = {
    props: ['progress',"legendValue","color","emptyColor","leftFormatted", "legendText"],
    methods: {
        timeAsHour: function (seconds) {
            if (seconds < 60) {
                return `${seconds}s`
            } else if (seconds < (60 * 60)) {
                return `${(seconds/60).toFixed(1)}m`
            } else {
                return `${(seconds/60/60).toFixed(1)}h`
            }
        },
        getPercent: function (project) {

            console.log("######", project)

            var an = project.wholeWorkAnalytics

            if (an.wholeSpent > an.wholeEstimated) {
                return 100
            }

            return an.wholeEstimated / an.wholeSpent * 100
        },
    },
    template: `
    <vue-ellipse-progress 
        v-bind:progress="progress"
        v-bind:legend-value="legendValue"
        :size="150"
        v-bind:color="color"
        v-bind:emptyColor="emptyColor"
        fontColor="#f5f3f1"
        >
        <span slot="legend-value" style="font-family: 'Fira Code',monospace">/{{ leftFormatted }}</span>
        <p slot="legend-caption" style="font-family: 'Fira Code',monospace; color: #f5f3f1">{{ legendText }}</p>
    </vue-ellipse-progress>
    `
}

const ProjectProgressBarKey = "rb-project-progress-bar"

var ProjectProgressBar = {
    props: ['project'],
    methods: {
        timeAsHour: function (seconds) {
            if (seconds < 60) {
                return `${seconds}s`
            } else if (seconds < (60 * 60)) {
                return `${(seconds / 60).toFixed(1)}m`
            } else {
                return `${(seconds / 60 / 60).toFixed(1)}h`
            }
        },
        getPercent: function (project) {

            var an = project.wholeWorkAnalytics

            console.log("######", an.wholeSpent, an.wholeEstimated)

            if (an.wholeSpent > an.wholeEstimated) {
                return 100
            }

            return an.wholeSpent / an.wholeEstimated  * 100
        },
        tasksProgress: function (project) {
            var all = (project.wholeWorkAnalytics.done + project.wholeWorkAnalytics.toDo)
            var r =  project.wholeWorkAnalytics.done / all * 100

            console.log("######", project.wholeWorkAnalytics.done, project.wholeWorkAnalytics.toDo)
            return r
        }
    },
    template:`
    <vs-card style="background-color: #313131" actionable class="cardx">
        <div slot="header">
            <rb-name-component
               title="Progress"
            ></rb-name-component>
        </div>
        <div style="display: flex;justify-content: center;align-items: center;">
            <rb-progress-bar
                v-bind:progress=" getPercent(project) " 
                v-bind:legendValue=" timeAsHour(project.wholeWorkAnalytics.wholeSpent) "
                color="rgba(14, 73, 181, 1)"
                emptyColor="rgba(14, 73, 181, 0.3)"
                v-bind:leftFormatted=" timeAsHour(project.wholeWorkAnalytics.wholeEstimated) "
                legendText="Spent/Estimated"
            ></rb-progress-bar>
            <rb-progress-bar
                v-bind:progress=" tasksProgress(project) " 
                v-bind:legendValue=" project.wholeWorkAnalytics.done"
                color="rgba(28, 153, 21, 1)"
                emptyColor="rgba(28, 153, 21, 0.3)"
                v-bind:leftFormatted=" project.wholeWorkAnalytics.done + project.wholeWorkAnalytics.toDo "
                legendText="Done/All"
                style="margin-left: 24px"
            ></rb-progress-bar>
        </div>
    </vs-card>
    `
}

const ProjectInfoKey = "rb-project-info"

var ProjectInfo = {
    props: ['project'],
    template: `
    <vs-row vs-justify="center" style="padding-left: 22px">
        <vs-col type="flex" vs-justify="left" vs-align="left" vs-w="8" style="width: 70%">
            <vs-card style="background-color: #414141" actionable class="cardx">
                <rb-project-header
                    v-bind:project="project"
                ></rb-project-header>
                <vs-row vs-justify="flex-start" style="margin-top: 8px">
                    <vs-col type="flex" vs-justify="left" vs-align="left" vs-w="5.9">
                        <rb-chart-component
                            title="Tracked"
                            v-bind:canvasId="project.name + 'CANVAS'">
                        </rb-chart-component>
                    </vs-col>
                    <vs-col type="flex" vs-justify="right" vs-align="right" vs-w="5.9" style="margin-left: 8px">
                        <rb-chart-component
                            title="Activity"
                            v-bind:canvasId="project.name + 'CANVAS_TIME'"
                        ></rb-chart-component>
                    </vs-col>
                </vs-row>
                <vs-row vs-justify="center">
                    <vs-col type="flex" vs-justify="center" vs-align="center">
                        <rb-project-progress-bar
                            v-bind:project="project" 
                        ></rb-project-progress-bar>
                    </vs-col>
                </vs-row>
            </vs-card>
        </vs-col>
    </vs-row>
`
}