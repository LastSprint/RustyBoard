
const NameComponentKey = "rb-name-component"

var NameComponent = {
    props: ['title', 'avatar_src'],
    template: `
<!--    <vs-avatar size="large" :src="avatar_src" v-if="avatar_src != null"/>-->
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
            </vs-card>
        </vs-col>
    </vs-row>
`
}