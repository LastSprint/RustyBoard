
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
    props: ['progress',"legendValue","color","emptyColor","leftFormatted", "legendText", 'size'],
    template: `
    <vue-ellipse-progress 
        v-bind:progress="progress"
        v-bind:legend-value="legendValue"
        :size="size"
        v-bind:color="color"
        v-bind:emptyColor="emptyColor"
        fontColor="#f5f3f1"
        
        >
        <span slot="legend-value" style="font-family: 'Fira Code',monospace">{{ leftFormatted }}</span>
        <p slot="legend-caption" style="font-family: 'Fira Code',monospace; color: #f5f3f1">{{ legendText }}</p>
    </vue-ellipse-progress>
    `
}

const UserCardKey = "rb-user-card"

var UserCard = {
    props: ['userWork', 'prjAn'],
    methods: {
        progress: function (left, right) {
            var calc = left / right * 100
            if (isNaN(calc)) { return 0 }
            var r = parseFloat(calc.toFixed(1))
            return r
        }
    },
    template:`
    <vs-card style="background-color: #313131" actionable class="cardx">
        <div slot="header">
            <vs-row vs-justify="flex-start">
                <vs-avatar
                    v-bind:src="userWork.user.imgUrl"
                    size="large" 
                ></vs-avatar>
                <rb-name-component
                    v-bind:title="userWork.user.name"
                ></rb-name-component>
            </vs-row>
        </div>
        <vs-row>
            <rb-chart-component
                title="Activity"
                :canvasId="userWork.user.name+'USER_CANVAS'"
            ></rb-chart-component>
        </vs-row>
        <vs-row style="display: flex;justify-content: center;align-items: center;">
            <vs-col type="flex" vs-justify="left" vs-align="left" vs-w="4.5">
                <vs-row>
                    <rb-progress-bar
                        v-bind:progress=" progress(userWork.workAnalytics.taskSpent, prjAn.taskSpent) " 
                        color="rgba(14, 73, 181, 1)"
                        emptyColor="rgba(14, 73, 181, 0.3)"
                        legendText="In Tasks"
                        leftFormatted="%"
                        style="margin-left: 24px"
                        :size="100"
                    ></rb-progress-bar>
                </vs-row>
                <vs-row style="margin-top: 16px">
                    <rb-progress-bar
                        v-bind:progress=" progress(userWork.workAnalytics.bugSpent, prjAn.bugSpent) " 
                        color="rgba(202, 62, 71, 1)"
                        emptyColor="rgba(202, 62, 71, 0.3)"
                        leftFormatted="%"
                        legendText="In Bugs"
                        style="margin-left: 24px"
                        :size="100"
                    ></rb-progress-bar>
                </vs-row>
            </vs-col>
            <vs-col type="flex" vs-justify="left" vs-align="left" vs-w="4.5">
                <vs-row>
                    <rb-progress-bar
                        v-bind:progress=" progress(userWork.workAnalytics.serviceSpent, prjAn.serviceSpent) " 
                        color="rgba(221, 221, 221, 1)"
                        emptyColor="rgba(221, 221, 221, 0.3)"
                        leftFormatted="%"
                        legendText="In S-Task"
                        style="margin-left: 24px"
                        :size="100"
                    ></rb-progress-bar>
                </vs-row>
                <vs-row style="margin-top: 16px">
                    <rb-progress-bar
                        v-bind:progress=" progress(userWork.workAnalytics.testSpent, prjAn.testSpent) " 
                        color="rgba(22, 199, 154, 1)"
                        emptyColor="rgba(22, 199, 154, 0.3)"
                        leftFormatted="%"
                        legendText="In Tasks"
                        style="margin-left: 24px"
                        :size="100"
                    ></rb-progress-bar>
                </vs-row>
            </vs-col>
        </vs-row>        
    </vs-card>
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
            if (an.wholeSpent > an.wholeEstimated) {
                return 100
            }

            return an.wholeSpent / an.wholeEstimated  * 100
        },
        tasksProgress: function (project) {
            var all = (project.wholeWorkAnalytics.done + project.wholeWorkAnalytics.toDo)
            return project.wholeWorkAnalytics.done / all * 100
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
                v-bind:leftFormatted=" '/'+timeAsHour(project.wholeWorkAnalytics.wholeEstimated) "
                legendText="Spent/Estimated"
                :size="150"
            ></rb-progress-bar>
            <rb-progress-bar
                v-bind:progress=" tasksProgress(project) " 
                v-bind:legendValue=" project.wholeWorkAnalytics.done"
                color="rgba(28, 153, 21, 1)"
                emptyColor="rgba(28, 153, 21, 0.3)"
                v-bind:leftFormatted=" '/' + (project.wholeWorkAnalytics.done + project.wholeWorkAnalytics.toDo) "
                legendText="Done/All"
                style="margin-left: 24px"
                :size="150"
            ></rb-progress-bar>
        </div>
    </vs-card>
    `
}

const ProjectInfoKey = "rb-project-info"

const UserCardWrapperKey = "rb-user-card-wrapper"

var UserCardWrapper = {
    props: ['userWork', 'prjAn'],
    template:`
    <vs-col 
        type="flex" 
        vs-justify="left" 
        vs-align="left" 
        vs-w="3.85"
    >
        <rb-user-card
            v-bind:userWork="userWork"
            v-bind:prjAn="prjAn"
        ></rb-user-card>
    </vs-col>
    `
}

var ProjectInfo = {
    props: ['project'],
    template: `
    <vs-row vs-justify="center" style="padding-left: 22px">
        <vs-col type="flex" vs-justify="left" vs-align="left" vs-w="8.5" style="width: 90%">
            <vs-card style="background-color: #414141" actionable class="cardx">
                <rb-project-header
                    style="font-size: 24px"
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
                <vs-row>
                    <vs-col type="flex" vs-justify="center" vs-align="center">
                        <vs-collapse open-hover >
                            <vs-collapse-item style="color: #f5f3f1" icon-arrow="⬇">
                                <div slot="header">
                                    <rb-name-component
                                        style="font-size: 24px"
                                        title="Who Works"
                                    ></rb-name-component>
                                </div>
                                <rb-user-card-wrapper 
                                    style="margin-left: 16px"
                                    v-for="it in project.whoWorks"
                                    v-bind:key="it.user.name + \`ROOT_USER_ID\`"
                                    v-bind:userWork="it"
                                    v-bind:prjAn="project.wholeWorkAnalytics"
                                ></rb-user-card-wrapper>
                            </vs-collapse-item>
                        </vs-collapse>
                    </vs-col>
                </vs-row>
            </vs-card>
        </vs-col>
    </vs-row>
`
}